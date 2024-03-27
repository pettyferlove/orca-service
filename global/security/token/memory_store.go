package token

import (
	"errors"
	"github.com/google/uuid"
	"orca-service/global/security"
	"sync"
	"time"
)

type MemoryStore struct {
	data sync.Map
}

type tokenData struct {
	user             security.UserDetail
	lastActivityTime time.Time
}

func NewMemoryStore() *MemoryStore {
	store := &MemoryStore{}

	// 设置定时任务，每分钟检查一次过期的令牌
	go store.startCleaningJob()

	return store
}

func (m *MemoryStore) CreateAccessToken(user security.UserDetail) (string, error) {
	uuidObj, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	token := uuidObj.String()
	m.data.Store(token, &tokenData{user: user, lastActivityTime: time.Now()})
	return token, nil
}

func (m *MemoryStore) RefreshAccessToken(token string) (string, error) {
	value, ok := m.data.Load(token)
	if !ok {
		return "", errors.New("invalid token")
	}
	value.(*tokenData).lastActivityTime = time.Now()
	return token, nil
}

func (m *MemoryStore) RemoveAccessToken(user security.UserDetail) error {
	m.data.Range(func(key, value interface{}) bool {
		data := value.(*tokenData)
		if data.user.GetId() == user.GetId() {
			m.data.Delete(key)
		}
		return true
	})
	return nil
}

func (m *MemoryStore) VerifyAccessToken(token string) (*security.UserDetail, error) {
	value, ok := m.data.Load(token)
	if !ok {
		return nil, errors.New("invalid token")
	}

	data := value.(*tokenData)
	data.lastActivityTime = time.Now() // 更新最后活跃时间
	return &data.user, nil
}

func (m *MemoryStore) startCleaningJob() {
	ticker := time.NewTicker(time.Minute)
	for range ticker.C {
		m.cleanExpiredTokens()
	}
}

func (m *MemoryStore) cleanExpiredTokens() {
	now := time.Now()
	m.data.Range(func(key, value interface{}) bool {
		data := value.(*tokenData)
		if now.Sub(data.lastActivityTime) > time.Hour {
			m.data.Delete(key)
		}
		return true
	})
}

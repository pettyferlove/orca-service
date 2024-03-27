package token

import (
	"github.com/stretchr/testify/assert"
	"orca-service/global/security"
	"testing"
	"time"
)

func TestMemoryStore(t *testing.T) {
	store := NewMemoryStore()
	user := security.UserDetail{
		Id:       "1",
		Username: "testUser",
		Email:    "test@example.com",
	}

	// 测试创建访问令牌
	token, err := store.CreateAccessToken(user)
	assert.Nil(t, err)
	assert.NotEmpty(t, token)

	// 测试验证访问令牌
	userDetail, err := store.VerifyAccessToken(token)
	assert.Nil(t, err)
	assert.Equal(t, user.GetId(), userDetail.GetId())

	// 测试刷新访问令牌
	newToken, err := store.RefreshAccessToken(token)
	assert.Nil(t, err)
	assert.Equal(t, token, newToken)

	// 测试移除访问令牌
	err = store.RemoveAccessToken(user)
	assert.Nil(t, err)

	// 验证移除后的访问令牌
	userDetail, err = store.VerifyAccessToken(token)
	assert.NotNil(t, err)
	assert.Nil(t, userDetail)
}

func TestMemoryStore_CleaningJob(t *testing.T) {
	store := NewMemoryStore()
	user := security.UserDetail{
		Id:       "1",
		Username: "testUser",
		Email:    "test@example.com",
	}

	token, err := store.CreateAccessToken(user)
	assert.Nil(t, err)
	assert.NotEmpty(t, token)

	// 手动设置令牌的最后活跃时间为2小时前，使其过期
	value, _ := store.data.Load(token)
	value.(*tokenData).lastActivityTime = time.Now().Add(-2 * time.Hour)

	// 执行清理过期令牌的任务
	store.cleanExpiredTokens()

	// 验证过期令牌已被清理
	_, ok := store.data.Load(token)
	assert.False(t, ok)
}

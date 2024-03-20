package token

import "orca-service/global/security/model"

type MemoryStore struct {
}

func (m MemoryStore) CreateAccessToken(user model.UserDetail) (string, error) {
	return "", nil
}

func (m MemoryStore) RefreshAccessToken(token string) (string, error) {
	return "", nil
}

func (m MemoryStore) RemoveAccessToken(user model.UserDetail) error {
	return nil
}

func (m MemoryStore) VerifyAccessToken(token string) (*model.UserDetail, error) {
	return nil, nil
}

package token

import "orca-service/global/security/model"

type RedisStore struct {
}

func (r RedisStore) CreateAccessToken(user model.UserDetail) (string, error) {
	return "", nil
}

func (r RedisStore) RefreshAccessToken(token string) (string, error) {
	return "", nil
}

func (r RedisStore) RemoveAccessToken(user model.UserDetail) error {
	return nil
}

func (r RedisStore) VerifyAccessToken(token string) (*model.UserDetail, error) {
	return nil, nil
}

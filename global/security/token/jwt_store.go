package token

import "orca-service/global/security/model"

type JwtStore struct {
}

func (j JwtStore) CreateAccessToken(user model.UserDetail) (string, error) {
	return "", nil
}

func (j JwtStore) RefreshAccessToken(token string) (string, error) {
	return "", nil
}

func (j JwtStore) RemoveAccessToken(user model.UserDetail) error {
	return nil
}

func (j JwtStore) VerifyAccessToken(token string) (*model.UserDetail, error) {
	return nil, nil
}

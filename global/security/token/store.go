package token

import (
	"orca-service/global/security"
)

// AllowMultiPoint 变量，是否允许多点登录
var AllowMultiPoint bool

// Store token store
type Store interface {

	// CreateAccessToken create access token
	CreateAccessToken(user security.UserDetail) (string, error)

	// RefreshAccessToken refresh access token
	RefreshAccessToken(token string) (string, error)

	// RemoveAccessToken remove access token
	RemoveAccessToken(user security.UserDetail) error

	// VerifyAccessToken verify access token
	VerifyAccessToken(token string) (*security.UserDetail, error)
}

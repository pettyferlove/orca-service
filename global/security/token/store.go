package token

import (
	"orca-service/global/security"
)

// Store token store
type Store interface {

	// CreateAccessToken 创建访问令牌
	CreateAccessToken(user security.UserDetail) (string, error)

	// RefreshAccessToken 刷新访问令牌
	RefreshAccessToken(token string) (string, error)

	// RemoveAccessToken 移除访问令牌
	RemoveAccessToken(user security.UserDetail) error

	// VerifyAccessToken 验证访问令牌
	VerifyAccessToken(token string) (*security.UserDetail, error)
}

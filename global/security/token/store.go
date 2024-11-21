package token

import (
	"orca-service/global/security"
	"sync/atomic"
)

var defaultStore atomic.Value

// Store token store
type Store interface {

	// CreateAccessToken 创建访问令牌
	CreateAccessToken(user security.UserDetail) (string, error)

	// RefreshAccessToken 刷新访问令牌
	RefreshAccessToken(token string) (string, error)

	// VerifyAccessToken 验证访问令牌
	VerifyAccessToken(token string) (*security.UserDetail, error)

	// RemoveAccessTokenByUser 移除指定用户的所有访问令牌
	RemoveAccessTokenByUser(user security.UserDetail) error

	// RemoveAccessTokenByToken 移除指定访问令牌
	RemoveAccessTokenByToken(token string) error

	// RemoveAllAccessToken 移除所有访问令牌
	RemoveAllAccessToken() error
}

// SetStore 设置默认的token store
func SetStore(store Store) {
	defaultStore.Store(store)
}

// GetStore 获取默认的token store
func GetStore() Store {
	return defaultStore.Load().(Store)
}

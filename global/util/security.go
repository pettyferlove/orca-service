package util

import (
	"context"
	"errors"
	"orca-service/global/security"
)

const (
	UserDetailKey  string = "orca/user_detail"
	AccessTokenKey string = "orca/access_token"
)

// WithContext 将JWTClaims保存到context中
func WithContext(ctx context.Context, user *security.UserDetail) context.Context {
	return context.WithValue(ctx, UserDetailKey, user)
}

// GetAccount 从上下文中获取用户信息
func GetAccount(ctx context.Context) (security.UserDetail, error) {
	if v := ctx.Value(UserDetailKey); v != nil {
		return v.(security.UserDetail), nil
	}
	return security.UserDetail{}, errors.New("user is not logged in")
}

// GetAccountId 从上下文中获取用户ID
func GetAccountId(ctx context.Context) string {
	if v := ctx.Value(UserDetailKey); v != nil {
		return v.(*security.UserDetail).Id
	}
	return "anonymous"
}

// GetRoles 从上下文中获取用户角色
func GetRoles(ctx context.Context) []string {
	if v := ctx.Value(UserDetailKey); v != nil {
		return v.(*security.UserDetail).Roles
	}
	return []string{}
}

// GetPermissions 从上下文中获取用户权限
func GetPermissions(ctx context.Context) []string {
	if v := ctx.Value(UserDetailKey); v != nil {
		return v.(*security.UserDetail).Permissions
	}
	return []string{}
}

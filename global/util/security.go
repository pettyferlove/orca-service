package util

import (
	"context"
	"orca-service/global/security"
	"orca-service/global/security/model"
)

type key string

const ContentKey key = "original_claims"

// WithContext 将JWTClaims保存到context中
func WithContext(ctx context.Context, claims *security.JWTClaims) context.Context {
	return context.WithValue(ctx, ContentKey, claims)
}

// GetAccount 从上下文中获取用户信息
func GetAccount(ctx context.Context) *model.UserDetail {
	if v := ctx.Value(ContentKey); v != nil {
		return &v.(*security.JWTClaims).UserDetail
	}
	return nil
}

// GetAccountId 从上下文中获取用户ID
func GetAccountId(ctx context.Context) string {
	if v := ctx.Value(ContentKey); v != nil {
		return v.(*security.JWTClaims).UserDetail.Id
	}
	return "anonymous"
}

// GetRoles 从上下文中获取用户角色
func GetRoles(ctx context.Context) []string {
	if v := ctx.Value(ContentKey); v != nil {
		return v.(*security.JWTClaims).Roles
	}
	return []string{}
}

// GetPermissions 从上下文中获取用户权限
func GetPermissions(ctx context.Context) []string {
	if v := ctx.Value(ContentKey); v != nil {
		return v.(*security.JWTClaims).Permissions
	}
	return []string{}
}

package model

import "time"

// LoginAttempts 登录尝试
type LoginAttempts struct {
	Id                string
	Username          string
	LoginFail         int
	LastLoginFailTime time.Time
}

// LoginRequest 登录请求
type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// LoginResponse 登录响应
type LoginResponse struct {
	Token string `json:"token"`
	Type  string `json:"type,default=Bearer"`
}

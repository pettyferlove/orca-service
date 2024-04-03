package model

import "time"

type LoginAttempts struct {
	Id                string
	Username          string
	LoginFail         int
	LastLoginFailTime time.Time
}

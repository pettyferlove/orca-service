package model

import "time"

type LoginAttempts struct {
	Id                string
	Username          string
	LoginFailNum      int
	LastLoginFailTime time.Time
}

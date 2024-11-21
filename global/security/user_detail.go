package security

import "time"

type UserStatus int8

const (
	Normal UserStatus = iota
	Locked
	Disabled
	Expired
)

type UserDetail struct {
	Id          string     `json:"id"`
	Avatar      string     `json:"avatar"`
	Username    string     `json:"username"`
	Password    string     `json:"password"`
	Nickname    string     `json:"nickname"`
	Email       string     `json:"email"`
	Phone       string     `json:"phone"`
	Channel     string     `json:"channel"`
	Status      UserStatus `json:"status"`
	ClientIp    string     `json:"client_ip"`
	LoginTime   time.Time  `json:"login_time"`
	Roles       []string   `json:"roles"`
	Permissions []string   `json:"permissions"`
}

func (u UserDetail) GetId() string {
	return u.Id
}

func (u UserDetail) GetUsername() string {
	return u.Username
}

func (u UserDetail) IsAccountNonExpired() bool {
	return u.Status != Expired
}

func (u UserDetail) IsAccountNonLocked() bool {
	return u.Status != Locked
}

func (u UserDetail) IsCredentialsNonExpired() bool {
	return u.Status != Expired
}

func (u UserDetail) IsNormal() bool {
	return u.Status == Normal
}

func (u UserDetail) IsDisabled() bool {
	return u.Status == Disabled
}

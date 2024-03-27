package model

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
	Tenant      string     `json:"tenant"`
	Status      UserStatus `json:"status"`
	Roles       []string   `json:"roles"`
	Permissions []string   `json:"permissions"`
}

func (u UserDetail) GetId() string {
	return u.Id
}

func (u UserDetail) GetUsername() string {
	return u.Username
}

func (u UserDetail) GetPassword() string {
	return u.Password
}

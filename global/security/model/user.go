package model

type UserDetail struct {
	Id       string `json:"id"`
	Avatar   string `json:"avatar"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Channel  string `json:"channel"`
	Tenant   string `json:"tenant"`
	Status   int8   `json:"status"`
}

package model

type UserDetail struct {
	Id       string `json:"id"`
	UserName string `json:"user_name"`
	NickName string `json:"nick_name"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
}

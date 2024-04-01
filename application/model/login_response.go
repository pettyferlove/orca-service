package model

type LoginResponse struct {
	Token string `json:"token"`
	Type  string `json:"type,default=Bearer"`
}

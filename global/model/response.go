package model

type Response struct {
	Timestamp  int64       `json:"timestamp"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	Code       int         `json:"code,default=0"`
	Successful bool        `json:"successful,default=true"`
}

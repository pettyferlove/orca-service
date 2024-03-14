package entity

type ResponseEntity struct {
	Code       int         `json:"code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	Successful bool        `json:"successful"`
}

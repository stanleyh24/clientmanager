package models

type Router struct {
	Id       int64  `json:"id,omitempty"`
	Ip       string `json:"ip" `
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

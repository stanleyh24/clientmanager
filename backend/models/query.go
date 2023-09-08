package models

type QueryParams struct {
	Method   string
	Username string
	Password string
	Ip       string
	Resource string
	Body     map[string]string
}

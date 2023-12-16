package models

type QueryParams struct {
	Method   string
	Resource string
	Body     map[string]string
}

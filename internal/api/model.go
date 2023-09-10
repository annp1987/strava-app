package api

// ErrorResponse Struct ...
type ErrorResponse struct {
	Resource string `json:"resource"`
	Field    string `json:"field"`
	Code     string `json:"code"`
}

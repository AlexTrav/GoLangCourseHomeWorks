package dto

type ErrorResponse struct {
	Error string `json:"error" example:"not_found"`
}

type ValidationErrorResponse struct {
	Error  string            `json:"error" example:"validation"`
	Fields map[string]string `json:"fields"`
}

package rest

// ErrorResponse defines response with error message
type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

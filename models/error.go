package models

type ErrorIn struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

type ErrorResponse struct {
	Error *ErrorIn `json:"error,omitempty"`
}

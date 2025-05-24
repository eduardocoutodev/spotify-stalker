package dto

type ErrorResponse struct {
	Error ErrorBody `json:"error"`
}

type ErrorBody struct {
	Message string `json:"message"`
}

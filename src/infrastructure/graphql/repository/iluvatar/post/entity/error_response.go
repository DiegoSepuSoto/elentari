package entity

type ErrorResponse struct {
	Errors []*ServiceError `json:"errors"`
}

type ServiceError struct {
	Message string `json:"message"`
}

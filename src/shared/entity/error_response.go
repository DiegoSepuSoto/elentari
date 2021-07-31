package entity

type ErrorResponse struct {
	Errors []*IluvatarError `json:"errors"`
}

type IluvatarError struct {
	Message string `json:"message"`
}

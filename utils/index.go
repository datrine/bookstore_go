package utils

type ErrResponse struct {
	Message string
}

type OkResponse struct {
	Message string
	Data    interface{}
}

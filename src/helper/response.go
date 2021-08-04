package helper

import (
	"strings"
)

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type EmptyObj struct{}

func BuildResponse(data interface{}) *Response {
	return &Response{
		Status: true,
		Data:   data,
	}
}

func BuildErrorResponse(message string, err string) *Response {
	splittedError := strings.Split(err, "\n")
	return &Response{
		Status:  false,
		Message: message,
		Errors:  splittedError,
	}
}

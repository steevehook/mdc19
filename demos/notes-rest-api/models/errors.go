package models

import (
	"net/http"
)

const (
	HTTPErrorType           = "http_error"
	DataValidationErrorType = "data_validation"
	NotFoundErrorType       = "not_found"
)

type ErrorHandlerFunc = func(w http.ResponseWriter, err error)

type DataValidationError struct {
	Message     string                 `json:"message"`
	Constraints map[string]interface{} `json:"constraints"`
}

func (e DataValidationError) Error() string {
	return e.Message
}

type HTTPError struct {
	StatusCode int    `json:"-"`
	Type       string `json:"type"`
	Message    string `json:"message"`
}

func (e HTTPError) Error() string {
	return e.Message
}

type NotFoundError struct{}

func (e NotFoundError) Error() string {
	return "resource not found"
}

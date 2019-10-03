package main

import (
	"encoding/json"
	"net/http"
)

func toJSON(w http.ResponseWriter, code int, data interface{}) {
	w.WriteHeader(code)
	// handle errors, bustard
	_ = json.NewEncoder(w).Encode(data)
}

func SendError(w http.ResponseWriter, code int, e error) {
	switch err := e.(type) {
	case DataValidationError:
		err.Type = DataValidationType
		toJSON(w, code, err)
	case HTTPError:
		err.Type = HTTPErrorType
		toJSON(w, code, err)
	default:
		toJSON(w, code, HTTPError{
			Type:    HTTPErrorType,
			Message: "some http error",
		})
	}
}

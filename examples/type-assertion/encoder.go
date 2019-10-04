package main

import (
	"encoding/json"
	"net/http"
)

func SendJSON(w http.ResponseWriter, code int, data interface{}) {
	w.WriteHeader(code)
	// handle errors, bustard
	_ = json.NewEncoder(w).Encode(data)
}

func SendError(w http.ResponseWriter, code int, e error) {
	switch err := e.(type) {
	case DataValidationError:
		err.Type = DataValidationType
		SendJSON(w, code, err)
	case HTTPError:
		err.Type = HTTPErrorType
		SendJSON(w, code, err)
	default:
		SendJSON(w, code, HTTPError{
			Type:    HTTPErrorType,
			Message: "some http error",
		})
	}
}

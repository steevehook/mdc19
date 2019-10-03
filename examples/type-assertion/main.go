package main

import (
	"errors"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/http", func(w http.ResponseWriter, req *http.Request) {
		SendError(w, http.StatusInternalServerError, HTTPError{Message: "something bad happen"})
	})
	http.HandleFunc("/data", func(w http.ResponseWriter, req *http.Request) {
		SendError(w, http.StatusBadRequest, DataValidationError{
			Message: "something bad happen",
			Constraints: map[string]interface{}{
				"field1": "some error on field1",
			},
		})
	})
	http.HandleFunc("/error", func(w http.ResponseWriter, req *http.Request) {
		SendError(w, http.StatusInternalServerError, errors.New("test"))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

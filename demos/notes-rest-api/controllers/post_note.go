package controllers

import (
	"net/http"
)

type notesPoster interface {
	CreateNote() error
}

func createNote(service notesPoster) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	}
}

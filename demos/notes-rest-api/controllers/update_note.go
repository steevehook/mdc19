package controllers

import (
	"net/http"
)

type notesUpdater interface {
	UpdateNote() error
}

func updateNote(service notesUpdater) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	}
}

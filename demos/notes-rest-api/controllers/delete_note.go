package controllers

import (
	"net/http"
)

type notesDeleter interface {
	DeleteNoteByID() error
	DeleteNotes() error
}

func deleteNote(service notesDeleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(204)
	}
}

package controllers

import (
	"net/http"
)

type notesGetter interface {
	GetNotes() (int, error)
	GetNoteByID() (int, error)
}

func getNotes(service notesGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	}
}

func getNoteByID(service notesGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	}
}

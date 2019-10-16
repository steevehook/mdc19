package controllers

import (
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"

	"github.com/steevehook/mdc19/notes-rest-api/middleware"
	"github.com/steevehook/mdc19/notes-rest-api/models"
	"github.com/steevehook/mdc19/notes-rest-api/transport"
)

type NotesService interface {
	notesGetter
	notesPoster
	notesUpdater
	notesDeleter
}

type RouterConfig struct {
	NotesService NotesService
}

func New(cfg RouterConfig) *httprouter.Router {
	chain := alice.New(middleware.Recovery(transport.SendPanicError))
	router := httprouter.New()

	router.GET("/_health", wrap(chain.ThenFunc(health)))
	router.GET("/notes", wrap(chain.ThenFunc(getNotes(cfg.NotesService))))
	router.GET("/notes/:"+models.NoteIDParam, wrap(chain.ThenFunc(getNoteByID(cfg.NotesService))))
	router.POST("/notes", wrap(chain.ThenFunc(createNote(cfg.NotesService))))
	router.PUT("/notes/:"+models.NoteIDParam, wrap(chain.ThenFunc(updateNote(cfg.NotesService))))
	router.DELETE("/notes/:"+models.NoteIDParam, wrap(chain.ThenFunc(deleteNote(cfg.NotesService))))
	router.NotFound = NotFound()

	return router
}

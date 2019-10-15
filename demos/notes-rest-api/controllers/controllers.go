package controllers

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/steevehook/mdc19/dummy-notes-rest-api/logging"
)

type key string

const (
	paramsKey key = "params"
)

func wrap(h http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, paramsKey, ps)

		h.ServeHTTP(w, r.WithContext(ctx))
	}
}

func params(r *http.Request) httprouter.Params {
	ctx := r.Context()
	psCtx := ctx.Value(paramsKey)
	ps, ok := psCtx.(httprouter.Params)

	if !ok {
		logging.Logger.Error("could not extract params from context")
		return httprouter.Params{}
	}
	return ps
}

package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func (a *application) routes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	if a.debug {
		r.Use(middleware.Logger)
	}

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(a.appName))
	})

	r.Get("/comments", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("List of comments"))
	})

	return r
}

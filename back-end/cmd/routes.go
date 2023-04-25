package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Get("/media/*", func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/media/", http.FileServer(http.Dir("./media"))).ServeHTTP(w, r)
	})

	mux.Get("/", Home)

	handler := enableCORS(mux)
	return handler
}
package main

import (
	"net/http"
	"taimekalender/back-end/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Get("/rooms", handlers.GetRooms)
	mux.Get("/plants/{roomID}", handlers.GetPlants)

	handler := enableCORS(mux)
	return handler
}

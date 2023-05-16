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

	// Authentication
	mux.Get("/login", handlers.Authenticate)
	mux.Post("/login", handlers.Login)
	mux.Post("/logout", handlers.Logout)

	// Rooms
	mux.Get("/rooms", handlers.GetRooms)
	mux.Post("/rooms", handlers.AddRoom)
	mux.Put("/rooms", handlers.UpdateRoom)
	mux.Delete("/rooms", handlers.DeleteRoom)

	// Plants
	mux.Get("/plants/{roomID}", handlers.GetPlants)
	mux.Post("/plants", handlers.AddPlant)
	mux.Put("/plants", handlers.UpdatePlant)
	mux.Delete("/plants", handlers.DeletePlant)

	handler := enableCORS(mux)
	return handler
}

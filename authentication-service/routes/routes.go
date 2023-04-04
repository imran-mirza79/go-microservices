package routes

import (
	"authentication/internal/config"
	"net/http"
	"github.com/go-chi/cors"
	"github.com/go-chi/chi/v5"
)

func Routes(cfg config.Config) http.Handler {
	router := chi.NewRouter();

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Acceot", "Authorized", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:[]string{"Link"},
		AllowCredentials: true,
		MaxAge: 300,
	}))

	return router	
}
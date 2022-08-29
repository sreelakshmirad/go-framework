package core

import (
	"fmt"

	"framework_v1/core/middlewares"
	"framework_v1/internal/callbacks"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

// Routes defined in the application
// also added cors middle ware to allow OPTIONS request method from the browsers
// embedded s *Server
// @return : chi.Router
func (s *Server) InitRouter() chi.Router {
	fmt.Println("Inside the router")
	r := chi.NewRouter()
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allowed all origins
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)
	r.Use(middleware.Logger)
	// r.Use(middleware.CleanPath)

	r.Use(middlewares.CustomLogger)
	mapURLs(r)

	r.Route("/fruits", func(r chi.Router) { // app/fruits
		r.Get("/", callbacks.GetAllFruits)

		r.Route("/{fruitID}", func(r chi.Router) {
			r.Get("/", callbacks.GetFruitByID)
		})
	})

	return r
}

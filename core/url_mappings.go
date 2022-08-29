package core

import (
	"fmt"
	"framework_v1/internal/callbacks"

	"github.com/go-chi/chi"
)

func mapURLs(r *chi.Mux) {

	fmt.Println("Inside mapping")
	r.Route("/fruits", func(r chi.Router) { // app/fruits
		r.Get("/", callbacks.GetAllFruits)
		r.Route("/{fruitID}", func(r chi.Router) {
			r.Get("/", callbacks.GetFruitByID)
		})
	})

}

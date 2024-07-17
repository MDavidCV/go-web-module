package main

import (
	"fmt"
	"net/http"

	"github.com/MDavidCV/go-web-module/cmd/server/handler"
	"github.com/go-chi/chi/v5"
)

func main() {

	controller := handler.NewProductController("/Users/dcastrillonv/Documents/meli-boootcamp/go/go-web/go-web-module/docs/db/products.json")

	router := chi.NewRouter()

	router.Route("/products", func(r chi.Router) {
		r.Get("/", controller.GetProducts())
		r.Post("/", controller.CreateProduct())
		r.Get("/{id}", controller.GetProductById())
		r.Get("/search", controller.SearchProduct())
	})

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", router)
}

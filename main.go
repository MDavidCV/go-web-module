package main

import (
	"fmt"
	"net/http"

	"github.com/MDavidCV/go-web-module/controllers"
	"github.com/MDavidCV/go-web-module/models"
	"github.com/MDavidCV/go-web-module/utility"
	"github.com/go-chi/chi/v5"
)

var products []models.Product

func main() {
	products = utility.LoadProducts()
	controller := controllers.NewProductController(products)

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

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

	router.Get("/products", controller.GetProducts())
	router.Get("/products/{id}", controller.GetProductById())
	router.Get("/products/search", controller.ProductSearch())

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", router)
}

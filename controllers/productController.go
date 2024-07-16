package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/MDavidCV/go-web-module/dtos"
	"github.com/MDavidCV/go-web-module/models"
	"github.com/go-chi/chi/v5"
)

type productController struct {
	st map[int]models.Product
}

func NewProductController(st []models.Product) *productController {
	stMap := make(map[int]models.Product, len(st))
	for _, product := range st {
		stMap[product.Id] = product
	}

	return &productController{
		st: stMap,
	}
}

func (pc *productController) GetProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		products := make([]models.Product, 0, len(pc.st))
		for _, product := range pc.st {
			products = append(products, product)
		}

		code := http.StatusOK
		body := &dtos.Response{
			Code:  code,
			Data:  products,
			Error: "",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		json.NewEncoder(w).Encode(body)
	}
}

func (pc *productController) GetProductById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))

		if err != nil {
			code := http.StatusBadRequest
			body := &dtos.Response{
				Code:  code,
				Data:  nil,
				Error: "Invalid id",
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(code)
			json.NewEncoder(w).Encode(body)
			return
		}

		product, ok := pc.st[id]
		if !ok {
			code := http.StatusNotFound
			body := &dtos.Response{
				Code:  code,
				Data:  nil,
				Error: "Product not found",
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(code)
			json.NewEncoder(w).Encode(body)
			return
		}

		code := http.StatusOK
		body := &dtos.Response{
			Code:  code,
			Data:  product,
			Error: "",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		json.NewEncoder(w).Encode(body)
	}
}

func (pc *productController) ProductSearch() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		priceGt, err := strconv.ParseFloat(r.URL.Query().Get("priceGt"), 64)

		if err != nil {
			code := http.StatusBadRequest
			body := &dtos.Response{
				Code:  code,
				Data:  nil,
				Error: "Invalid priceGt",
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(code)
			json.NewEncoder(w).Encode(body)
			return
		}

		products := []models.Product{}
		for _, product := range pc.st {
			if product.Price > priceGt {
				products = append(products, product)
			}
		}

		code := http.StatusOK
		body := &dtos.Response{
			Code:  code,
			Data:  products,
			Error: "",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		json.NewEncoder(w).Encode(body)
	}
}

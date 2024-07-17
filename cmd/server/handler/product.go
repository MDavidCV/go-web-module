package handler

import (
	"encoding/json"
	"net/http"

	"github.com/MDavidCV/go-web-module/internal/product"
	"github.com/MDavidCV/go-web-module/utility"
	"github.com/go-chi/chi/v5"
)

type ProductController interface {
	GetProducts() http.HandlerFunc
	GetProductById() http.HandlerFunc
	SearchProduct() http.HandlerFunc
	CreateProduct() http.HandlerFunc
	UpdateProduct() http.HandlerFunc
	UpdatePatchProduct() http.HandlerFunc
}

type productController struct {
	service product.ServiceProduct
}

func NewProductController(dataPath string) *productController {
	s := product.NewServiceProduct(dataPath)

	return &productController{
		service: s,
	}
}

func (pc *productController) GetProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		products, err := pc.service.GetProducts()

		if err != nil {
			response := utility.NewErrorResponse(err)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(response.Code)
			json.NewEncoder(w).Encode(response)
			return
		}

		response := utility.NewSuccessResponse(products)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.Code)
		json.NewEncoder(w).Encode(response)
	}
}

func (pc *productController) GetProductById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		product, err := pc.service.GetProductById(chi.URLParam(r, "id"))

		if err != nil {
			response := utility.NewErrorResponse(err)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(response.Code)
			json.NewEncoder(w).Encode(response)
			return
		}

		response := utility.NewSuccessResponse(product)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.Code)
		json.NewEncoder(w).Encode(response)
	}
}

func (pc *productController) SearchProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		productsFiltered, err := pc.service.SearchProduct(r.URL.Query().Get("priceGt"))

		if err != nil {
			response := utility.NewErrorResponse(err)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(response.Code)
			json.NewEncoder(w).Encode(response)
			return
		}

		response := utility.NewSuccessResponse(productsFiltered)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.Code)
		json.NewEncoder(w).Encode(response)
	}
}

func (pc *productController) CreateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var reqBody utility.ProductRequest
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			code := http.StatusBadRequest
			body := &utility.Response{
				Code:  code,
				Data:  nil,
				Error: "Invalid request body",
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(code)
			json.NewEncoder(w).Encode(body)
			return
		}

		product, err := pc.service.CreateProduct(reqBody)

		if err != nil {
			response := utility.NewErrorResponse(err)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(response.Code)
			json.NewEncoder(w).Encode(response)
			return
		}

		response := utility.NewSuccessResponse(product)
		response.Code = http.StatusCreated

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.Code)
		json.NewEncoder(w).Encode(response)
	}
}

func (pc *productController) UpdateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var reqBody utility.ProductRequest
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			code := http.StatusBadRequest
			body := &utility.Response{
				Code:  code,
				Data:  nil,
				Error: "Invalid request body",
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(code)
			json.NewEncoder(w).Encode(body)
			return
		}

		product, err := pc.service.UpdateProduct(chi.URLParam(r, "id"), reqBody)

		if err != nil {
			response := utility.NewErrorResponse(err)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(response.Code)
			json.NewEncoder(w).Encode(response)
			return
		}

		response := utility.NewSuccessResponse(product)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.Code)
		json.NewEncoder(w).Encode(response)
	}
}
func (pc *productController) DeleteProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		err := pc.service.DeleteProduct(chi.URLParam(r, "id"))

		if err != nil {
			response := utility.NewErrorResponse(err)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(response.Code)
			json.NewEncoder(w).Encode(response)
			return
		}

		response := utility.NewSuccessResponse(nil)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.Code)
		json.NewEncoder(w).Encode(response)
	}
}

func (pc *productController) UpdatePatchProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var reqBody utility.ProductPatchRequest
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			code := http.StatusBadRequest
			body := &utility.Response{
				Code:  code,
				Data:  nil,
				Error: "Invalid request body",
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(code)
			json.NewEncoder(w).Encode(body)
			return
		}

		product, err := pc.service.UpdatePatchProduct(chi.URLParam(r, "id"), reqBody)

		if err != nil {
			response := utility.NewErrorResponse(err)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(response.Code)
			json.NewEncoder(w).Encode(response)
			return
		}

		response := utility.NewSuccessResponse(product)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.Code)
		json.NewEncoder(w).Encode(response)
	}
}

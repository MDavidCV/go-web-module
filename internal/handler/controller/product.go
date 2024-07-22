package controller

import (
	"encoding/json"
	"net/http"

	"github.com/MDavidCV/go-web-module/internal/domain"
	"github.com/MDavidCV/go-web-module/internal/service"
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
	service service.ServiceProduct
}

func (pc *productController) GetProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		products, err := pc.service.GetProducts()

		if err != nil {
			HandleResponse(w, utility.NewErrorResponse(err))
			return
		}

		HandleResponse(w, utility.NewSuccessResponse(products))
	}
}

func (pc *productController) GetProductById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		product, err := pc.service.GetProductById(chi.URLParam(r, "id"))

		if err != nil {
			HandleResponse(w, utility.NewErrorResponse(err))
			return
		}

		HandleResponse(w, utility.NewSuccessResponse(product))
	}
}

func (pc *productController) SearchProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		productsFiltered, err := pc.service.SearchProduct(r.URL.Query().Get("priceGt"))

		if err != nil {
			HandleResponse(w, utility.NewErrorResponse(err))
			return
		}

		HandleResponse(w, utility.NewSuccessResponse(productsFiltered))
	}
}

func (pc *productController) CreateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var reqBody utility.ProductRequest
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			err = utility.ErrInvalidRequestBody
			HandleResponse(w, utility.NewErrorResponse(err))
			return
		}

		product, err := pc.service.CreateProduct(reqBody)

		if err != nil {
			HandleResponse(w, utility.NewErrorResponse(err))
			return
		}

		response := utility.NewSuccessResponse(product)
		response.Code = http.StatusCreated
		HandleResponse(w, response)
	}
}

func (pc *productController) UpdateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var reqBody utility.ProductRequest
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			err = utility.ErrInvalidRequestBody
			HandleResponse(w, utility.NewErrorResponse(err))
			return
		}

		product, err := pc.service.UpdateProduct(chi.URLParam(r, "id"), reqBody)
		if err != nil {
			HandleResponse(w, utility.NewErrorResponse(err))
			return
		}

		HandleResponse(w, utility.NewSuccessResponse(product))
	}
}
func (pc *productController) DeleteProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		err := pc.service.DeleteProduct(chi.URLParam(r, "id"))
		if err != nil {
			HandleResponse(w, utility.NewErrorResponse(err))
			return
		}

		response := utility.NewSuccessResponse(nil)
		response.Code = http.StatusNoContent
		HandleResponse(w, response)
	}
}

func (pc *productController) UpdatePatchProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var reqBody utility.ProductPatchRequest
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			err = utility.ErrInvalidRequestBody
			HandleResponse(w, utility.NewErrorResponse(err))
			return
		}

		product, err := pc.service.UpdatePatchProduct(chi.URLParam(r, "id"), reqBody)
		if err != nil {
			HandleResponse(w, utility.NewErrorResponse(err))
			return
		}

		HandleResponse(w, utility.NewSuccessResponse(product))
	}
}
func (pc *productController) GetConsumerPrice() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("list")
		products, totalPrice, err := pc.service.GetConsumerPrice(query)

		if err != nil {
			HandleResponse(w, utility.NewErrorResponse(err))
			return
		}

		data := struct {
			Products   []domain.Product
			TotalPrice float64
		}{Products: products, TotalPrice: totalPrice}
		HandleResponse(w, utility.NewSuccessResponse(data))
	}
}

func NewProductController(service service.ServiceProduct) *productController {
	return &productController{
		service: service,
	}
}

func HandleResponse(w http.ResponseWriter, response utility.Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.Code)
	json.NewEncoder(w).Encode(response)
}

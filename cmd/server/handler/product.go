package handler

import (
	"encoding/json"
	"net/http"
	"os"

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
			handleResponse(w, utility.NewErrorResponse(err))
			return
		}

		handleResponse(w, utility.NewSuccessResponse(products))
	}
}

func (pc *productController) GetProductById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		product, err := pc.service.GetProductById(chi.URLParam(r, "id"))

		if err != nil {
			handleResponse(w, utility.NewErrorResponse(err))
			return
		}

		handleResponse(w, utility.NewSuccessResponse(product))
	}
}

func (pc *productController) SearchProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		productsFiltered, err := pc.service.SearchProduct(r.URL.Query().Get("priceGt"))

		if err != nil {
			handleResponse(w, utility.NewErrorResponse(err))
			return
		}

		handleResponse(w, utility.NewSuccessResponse(productsFiltered))
	}
}

func (pc *productController) CreateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if !checkAuthorization(r.Header.Get("token")) {
			handleResponse(w, utility.NewUnauthorizedResponse())
			return
		}

		var reqBody utility.ProductRequest
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			err = utility.ErrInvalidRequestBody
			handleResponse(w, utility.NewErrorResponse(err))
			return
		}

		product, err := pc.service.CreateProduct(reqBody)

		if err != nil {
			handleResponse(w, utility.NewErrorResponse(err))
			return
		}

		response := utility.NewSuccessResponse(product)
		response.Code = http.StatusCreated
		handleResponse(w, response)
	}
}

func (pc *productController) UpdateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if !checkAuthorization(r.Header.Get("token")) {
			handleResponse(w, utility.NewUnauthorizedResponse())
			return
		}

		var reqBody utility.ProductRequest
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			err = utility.ErrInvalidRequestBody
			handleResponse(w, utility.NewErrorResponse(err))
			return
		}

		product, err := pc.service.UpdateProduct(chi.URLParam(r, "id"), reqBody)
		if err != nil {
			handleResponse(w, utility.NewErrorResponse(err))
			return
		}

		handleResponse(w, utility.NewSuccessResponse(product))
	}
}
func (pc *productController) DeleteProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if !checkAuthorization(r.Header.Get("token")) {
			handleResponse(w, utility.NewUnauthorizedResponse())
			return
		}

		err := pc.service.DeleteProduct(chi.URLParam(r, "id"))
		if err != nil {
			handleResponse(w, utility.NewErrorResponse(err))
			return
		}

		handleResponse(w, utility.NewSuccessResponse(nil))
	}
}

func (pc *productController) UpdatePatchProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if !checkAuthorization(r.Header.Get("token")) {
			handleResponse(w, utility.NewUnauthorizedResponse())
			return
		}

		var reqBody utility.ProductPatchRequest
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			err = utility.ErrInvalidRequestBody
			handleResponse(w, utility.NewErrorResponse(err))
			return
		}

		product, err := pc.service.UpdatePatchProduct(chi.URLParam(r, "id"), reqBody)
		if err != nil {
			handleResponse(w, utility.NewErrorResponse(err))
			return
		}

		handleResponse(w, utility.NewSuccessResponse(product))
	}
}

func checkAuthorization(token string) bool {
	api_key := os.Getenv("API_KEY")
	return api_key == token
}

func handleResponse(w http.ResponseWriter, response utility.Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.Code)
	json.NewEncoder(w).Encode(response)
}

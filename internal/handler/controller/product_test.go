package controller_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/MDavidCV/go-web-module/internal/domain"
	"github.com/MDavidCV/go-web-module/internal/handler/controller"
	"github.com/MDavidCV/go-web-module/internal/handler/middleware"
	"github.com/MDavidCV/go-web-module/internal/repository"
	"github.com/MDavidCV/go-web-module/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
)

// Happy Path Tests ----------------

func TestProductsGet(t *testing.T) {
	t.Run("sucess should return a list of products", func(t *testing.T) {
		// Arrange
		mockSt := map[int]domain.Product{
			1: {
				Id:          1,
				Name:        "Product 1",
				Quantity:    10,
				CodeValue:   "12345",
				IsPublished: true,
				Expiration:  "2023-01-01",
				Price:       100.0,
			},
			2: {
				Id:          2,
				Name:        "Product 2",
				Quantity:    20,
				CodeValue:   "67890",
				IsPublished: false,
				Expiration:  "2023-01-02",
				Price:       200.0,
			},
		}
		mockRepository := repository.NewRepositoryProduct(mockSt, nil)
		service := service.NewServiceProduct(mockRepository)
		controller := controller.NewProductController(service)

		// Act
		r := httptest.NewRequest("GET", "/products", nil)
		w := httptest.NewRecorder()
		controller.GetProducts()(w, r)

		// Assert
		expectedCode := http.StatusOK
		expectedBody := `{"body":[{"id":1,"name":"Product 1","quantity":10,"code_value":"12345","is_published":true,"expiration":"2023-01-01","price":100},{"id":2,"name":"Product 2","quantity":20,"code_value":"67890","is_published":false,"expiration":"2023-01-02","price":200}], "code": 200, "error": ""}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		require.Equal(t, expectedCode, w.Code)
		require.JSONEq(t, expectedBody, w.Body.String())
		require.Equal(t, expectedHeader, w.Header())
	})
}

func TestProuductByIdGet(t *testing.T) {
	t.Run("sucess should return a product", func(t *testing.T) {
		// Arrange
		mockSt := map[int]domain.Product{
			1: {
				Id:          1,
				Name:        "Product 1",
				Quantity:    10,
				CodeValue:   "12345",
				IsPublished: true,
				Expiration:  "2023-01-01",
				Price:       100.0,
			},
			2: {
				Id:          2,
				Name:        "Product 2",
				Quantity:    20,
				CodeValue:   "67890",
				IsPublished: false,
				Expiration:  "2023-01-02",
				Price:       200.0,
			},
		}
		mockRepository := repository.NewRepositoryProduct(mockSt, nil)
		service := service.NewServiceProduct(mockRepository)
		controller := controller.NewProductController(service)

		// Act
		r := httptest.NewRequest("GET", "/products/1", nil)
		w := httptest.NewRecorder()

		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "1")
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, chiCtx))

		controller.GetProductById()(w, r)

		// Assert
		expectedCode := http.StatusOK
		expectedBody := `{"body":{"id":1,"name":"Product 1","quantity":10,"code_value":"12345","is_published":true,"expiration":"2023-01-01","price":100}, "code": 200, "error": ""}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		require.Equal(t, expectedCode, w.Code)
		require.JSONEq(t, expectedBody, w.Body.String())
		require.Equal(t, expectedHeader, w.Header())
	})
}

func TestCreateProduct(t *testing.T) {
	t.Run("sucess should create a product", func(t *testing.T) {
		// Arrange
		mockSt := map[int]domain.Product{}
		mockRepository := repository.NewRepositoryProduct(mockSt, nil)
		service := service.NewServiceProduct(mockRepository)
		controller := controller.NewProductController(service)

		product := `{"name": "test", "quantity": 23, "code_value": "testcode", "is_published": true, "expiration": "15/12/2021", "price": 99}`

		// Act
		r := httptest.NewRequest("POST", "/products", strings.NewReader(product))
		w := httptest.NewRecorder()
		controller.CreateProduct()(w, r)

		// Assert
		expectedCode := http.StatusCreated
		expectedBody := `{"body":{"name": "test", "quantity": 23, "code_value": "testcode", "is_published": true, "expiration": "15/12/2021", "price": 99, "id": 1}, "code": 201, "error": ""}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		require.Equal(t, expectedCode, w.Code)
		require.JSONEq(t, expectedBody, w.Body.String())
		require.Equal(t, expectedHeader, w.Header())
	})
}

func TestDeleteProduct(t *testing.T) {
	t.Run("sucess should delete a product", func(t *testing.T) {
		// Arrange
		// Arrange
		mockSt := map[int]domain.Product{
			1: {
				Id:          1,
				Name:        "Product 1",
				Quantity:    10,
				CodeValue:   "12345",
				IsPublished: true,
				Expiration:  "2023-01-01",
				Price:       100.0,
			},
		}
		mockRepository := repository.NewRepositoryProduct(mockSt, nil)
		service := service.NewServiceProduct(mockRepository)
		controller := controller.NewProductController(service)

		// Act
		r := httptest.NewRequest("DELETE", "/products/1", nil)
		w := httptest.NewRecorder()

		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "1")
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, chiCtx))

		controller.DeleteProduct()(w, r)

		// Assert
		expectedCode := http.StatusNoContent
		expectedBody := `{"body":null, "code": 204, "error": ""}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		require.Equal(t, expectedCode, w.Code)
		require.JSONEq(t, expectedBody, w.Body.String())
		require.Equal(t, expectedHeader, w.Header())
	})
}

// Error Path Tests ----------------

func TestBadGetProductById(t *testing.T) {
	t.Run("should return an error when the id is not a number", func(t *testing.T) {
		// Arrange
		mockSt := map[int]domain.Product{
			1: {
				Id:          1,
				Name:        "Product 1",
				Quantity:    10,
				CodeValue:   "12345",
				IsPublished: true,
				Expiration:  "2023-01-01",
				Price:       100.0,
			},
		}
		mockRepository := repository.NewRepositoryProduct(mockSt, nil)
		service := service.NewServiceProduct(mockRepository)
		controller := controller.NewProductController(service)

		// Act
		r := httptest.NewRequest("GET", "/products/test", nil)
		w := httptest.NewRecorder()

		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "test")
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, chiCtx))

		controller.GetProductById()(w, r)

		// Assert
		expectedCode := http.StatusBadRequest
		expectedBody := `{"body":null, "code": 400, "error": "invalid id"}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		require.Equal(t, expectedCode, w.Code)
		require.JSONEq(t, expectedBody, w.Body.String())
		require.Equal(t, expectedHeader, w.Header())
	})
}

func TestGetUnexistentProductById(t *testing.T) {
	t.Run("should return an error when the product does not exist", func(t *testing.T) {
		// Arrange
		mockSt := map[int]domain.Product{
			1: {
				Id:          1,
				Name:        "Product 1",
				Quantity:    10,
				CodeValue:   "12345",
				IsPublished: true,
				Expiration:  "2023-01-01",
				Price:       100.0,
			},
		}
		mockRepository := repository.NewRepositoryProduct(mockSt, nil)
		service := service.NewServiceProduct(mockRepository)
		controller := controller.NewProductController(service)

		// Act
		r := httptest.NewRequest("GET", "/products/2", nil)
		w := httptest.NewRecorder()

		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "2")
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, chiCtx))

		controller.GetProductById()(w, r)

		// Assert
		expectedCode := http.StatusNotFound
		expectedBody := `{"body":null, "code": 404, "error": "product not found"}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		require.Equal(t, expectedCode, w.Code)
		require.JSONEq(t, expectedBody, w.Body.String())
		require.Equal(t, expectedHeader, w.Header())
	})
}

func TestPatchUnexistentProduct(t *testing.T) {
	t.Run("should return an error when the product does not exist", func(t *testing.T) {
		// Arrange
		mockSt := map[int]domain.Product{
			1: {
				Id:          1,
				Name:        "Product 1",
				Quantity:    10,
				CodeValue:   "12345",
				IsPublished: true,
				Expiration:  "2023-01-01",
				Price:       100.0,
			},
		}
		mockRepository := repository.NewRepositoryProduct(mockSt, nil)
		service := service.NewServiceProduct(mockRepository)
		controller := controller.NewProductController(service)

		productPatch := `{"name": "test patch"}`

		// Act
		r := httptest.NewRequest("GET", "/products/2", strings.NewReader(productPatch))
		w := httptest.NewRecorder()

		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "2")
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, chiCtx))

		controller.UpdatePatchProduct()(w, r)

		// Assert
		expectedCode := http.StatusNotFound
		expectedBody := `{"body":null, "code": 404, "error": "product not found"}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		require.Equal(t, expectedCode, w.Code)
		require.JSONEq(t, expectedBody, w.Body.String())
		require.Equal(t, expectedHeader, w.Header())
	})
}

func TestDeleteUnexistentProduct(t *testing.T) {
	t.Run("should return an error when the product does not exist", func(t *testing.T) {
		// Arrange
		mockSt := map[int]domain.Product{
			1: {
				Id:          1,
				Name:        "Product 1",
				Quantity:    10,
				CodeValue:   "12345",
				IsPublished: true,
				Expiration:  "2023-01-01",
				Price:       100.0,
			},
		}
		mockRepository := repository.NewRepositoryProduct(mockSt, nil)
		service := service.NewServiceProduct(mockRepository)
		controller := controller.NewProductController(service)

		// Act
		r := httptest.NewRequest("DELETE", "/products/2", nil)
		w := httptest.NewRecorder()

		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "2")
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, chiCtx))

		controller.DeleteProduct()(w, r)

		// Assert
		expectedCode := http.StatusNotFound
		expectedBody := `{"body":null, "code": 404, "error": "product not found"}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		require.Equal(t, expectedCode, w.Code)
		require.JSONEq(t, expectedBody, w.Body.String())
		require.Equal(t, expectedHeader, w.Header())
	})
}

func TestUnAuthDeleteProduct(t *testing.T) {
	t.Run("should return an error unauth", func(t *testing.T) {
		// Arrange
		mockSt := map[int]domain.Product{
			1: {
				Id:          1,
				Name:        "Product 1",
				Quantity:    10,
				CodeValue:   "12345",
				IsPublished: true,
				Expiration:  "2023-01-01",
				Price:       100.0,
			},
		}
		mockRepository := repository.NewRepositoryProduct(mockSt, nil)
		service := service.NewServiceProduct(mockRepository)
		controller := controller.NewProductController(service)

		router := chi.NewRouter()
		router.Use(middleware.AuthValidationMid)
		router.Delete("/products/{id}", controller.DeleteProduct())

		// Act
		r := httptest.NewRequest("DELETE", "/products/1", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, r)

		// Assert
		expectedCode := http.StatusUnauthorized
		expectedBody := `{"body":null, "code": 401, "error": "Unauthorized - Invalid Token"}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		require.Equal(t, expectedCode, w.Code)
		require.JSONEq(t, expectedBody, w.Body.String())
		require.Equal(t, expectedHeader, w.Header())
	})
}

func TestUnAuthPostProduct(t *testing.T) {
	t.Run("should return an error unauth", func(t *testing.T) {
		// Arrange
		mockSt := map[int]domain.Product{
			1: {
				Id:          1,
				Name:        "Product 1",
				Quantity:    10,
				CodeValue:   "12345",
				IsPublished: true,
				Expiration:  "2023-01-01",
				Price:       100.0,
			},
		}
		mockRepository := repository.NewRepositoryProduct(mockSt, nil)
		service := service.NewServiceProduct(mockRepository)
		controller := controller.NewProductController(service)

		product := `{"name": "test", "quantity": 23, "code_value": "testcode", "is_published": true, "expiration": "15/12/2021", "price": 99}`

		router := chi.NewRouter()
		router.Use(middleware.AuthValidationMid)
		router.Delete("/products/{id}", controller.DeleteProduct())

		// Act
		r := httptest.NewRequest("POST", "/products", strings.NewReader(product))
		w := httptest.NewRecorder()

		router.ServeHTTP(w, r)

		// Assert
		expectedCode := http.StatusUnauthorized
		expectedBody := `{"body":null, "code": 401, "error": "Unauthorized - Invalid Token"}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		require.Equal(t, expectedCode, w.Code)
		require.JSONEq(t, expectedBody, w.Body.String())
		require.Equal(t, expectedHeader, w.Header())
	})
}

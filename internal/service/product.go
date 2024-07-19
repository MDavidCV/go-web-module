package service

import (
	"strconv"
	"strings"

	"github.com/MDavidCV/go-web-module/internal/domain"
	"github.com/MDavidCV/go-web-module/internal/repository"
	"github.com/MDavidCV/go-web-module/utility"
)

type ServiceProduct interface {
	GetProducts() ([]domain.Product, error)
	GetProductById(pathVariable string) (domain.Product, error)
	SearchProduct(query string) ([]domain.Product, error)
	CreateProduct(product utility.ProductRequest) (domain.Product, error)
	UpdateProduct(pathVariable string, product utility.ProductRequest) (domain.Product, error)
	DeleteProduct(pathVariable string) error
	UpdatePatchProduct(pathVariable string, product utility.ProductPatchRequest) (domain.Product, error)
	GetConsumerPrice(query string) ([]domain.Product, float64, error)
}

type serviceProduct struct {
	repository repository.RepositoryProduct
}

func (sp *serviceProduct) GetProducts() ([]domain.Product, error) {
	return sp.repository.GetProducts()
}

func (sp *serviceProduct) GetProductById(pathVariable string) (domain.Product, error) {

	id, err := strconv.Atoi(pathVariable)

	if err != nil {
		return domain.Product{}, utility.ErrInvalidId
	}

	product, err := sp.repository.GetProductById(id)

	if err != nil {
		return domain.Product{}, err
	}

	return product, nil
}

func (sp *serviceProduct) SearchProduct(query string) ([]domain.Product, error) {
	priceGt, err := strconv.ParseFloat(query, 64)
	if err != nil {
		return nil, utility.ErrInvalidQuery
	}

	products, err := sp.repository.GetProducts()

	if err != nil {
		return nil, err
	}

	var productsFiltered []domain.Product
	for _, product := range products {
		if product.Price > priceGt {
			productsFiltered = append(productsFiltered, product)
		}
	}

	return productsFiltered, nil
}

func (sp *serviceProduct) CreateProduct(reqProduct utility.ProductRequest) (domain.Product, error) {
	products, err := sp.repository.GetProducts()

	if err != nil {
		return domain.Product{}, err
	}

	switch {
	case !reqProduct.VerifyNonZeroValues():
		return domain.Product{}, utility.ErrInvalidValues
	case !reqProduct.VerifyUniqueCodeValue(products):
		return domain.Product{}, utility.ErrUniqueCodeValue
	case !reqProduct.VerifyExpirationDate():
		return domain.Product{}, utility.ErrInvalidDate
	}

	return sp.repository.CreateProduct(reqProduct)
}

func (sp *serviceProduct) UpdateProduct(pathVariable string, reqProduct utility.ProductRequest) (domain.Product, error) {

	id, err := strconv.Atoi(pathVariable)
	if err != nil {
		return domain.Product{}, utility.ErrInvalidId
	}

	products, err := sp.repository.GetProducts()
	if err != nil {
		return domain.Product{}, err
	}

	switch {
	case !reqProduct.VerifyNonZeroValues():
		return domain.Product{}, utility.ErrInvalidValues
	case !reqProduct.VerifyUniqueCodeValue(products):
		return domain.Product{}, utility.ErrUniqueCodeValue
	case !reqProduct.VerifyExpirationDate():
		return domain.Product{}, utility.ErrInvalidDate
	}

	return sp.repository.UpdateProduct(id, reqProduct)
}

func (sp *serviceProduct) DeleteProduct(pathVariable string) error {
	id, err := strconv.Atoi(pathVariable)
	if err != nil {
		return utility.ErrInvalidId
	}

	return sp.repository.DeleteProduct(id)
}

func (sp *serviceProduct) UpdatePatchProduct(pathVariable string, reqProduct utility.ProductPatchRequest) (domain.Product, error) {
	id, err := strconv.Atoi(pathVariable)
	if err != nil {
		return domain.Product{}, utility.ErrInvalidId
	}

	products, err := sp.repository.GetProducts()
	if err != nil {
		return domain.Product{}, err
	}

	switch {
	case !reqProduct.VerifyNonZeroValues():
		return domain.Product{}, utility.ErrInvalidValues
	case !reqProduct.VerifyUniqueCodeValue(products):
		return domain.Product{}, utility.ErrUniqueCodeValue
	case !reqProduct.VerifyExpirationDate():
		return domain.Product{}, utility.ErrInvalidDate
	}

	return sp.repository.UpdatePatchProduct(id, reqProduct)
}

func (sp *serviceProduct) GetConsumerPrice(query string) ([]domain.Product, float64, error) {

	var products []domain.Product
	var totalPrice float64
	var totalProducts int

	if query == "" {
		var err error
		products, err = sp.repository.GetProducts()

		if err != nil {
			return nil, 0, err
		}

		for _, product := range products {
			totalPrice += product.Price
			totalProducts++
		}

	} else {
		rawValues := strings.Trim(query, "[]")
		values := strings.Split(rawValues, ",")
		uniqueProductsIds := utility.CountValues(values)

		products = []domain.Product{}
		for key, value := range uniqueProductsIds {
			id, err := strconv.Atoi(key)
			if err != nil {
				return nil, 0, utility.ErrInvalidQuery
			}

			product, err := sp.repository.GetProductById(id)
			if err != nil {
				return nil, 0, err
			}

			if value > product.Quantity {
				return nil, 0, utility.ErrInvalidQuery
			}
			if !product.IsPublished {
				return nil, 0, utility.ErrInvalidQuery
			}

			products = append(products, product)
			totalPrice += product.Price * float64(value)
			totalProducts += value
		}
	}

	switch {
	case totalProducts <= 10:
		totalPrice = totalPrice * 1.21
	case totalProducts > 10 && totalProducts <= 20:
		totalPrice = totalPrice * 1.17
	case totalProducts > 20:
		totalPrice = totalPrice * 1.15
	}

	return products, totalPrice, nil
}

func NewServiceProduct(repository repository.RepositoryProduct) *serviceProduct {
	return &serviceProduct{
		repository: repository,
	}
}

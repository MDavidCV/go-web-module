package product

import (
	"strconv"

	"github.com/MDavidCV/go-web-module/internal/domain"
	"github.com/MDavidCV/go-web-module/utility"
)

type ServiceProduct interface {
	GetProducts() ([]domain.Product, error)
	GetProductById(pathVariable string) (domain.Product, error)
	SearchProduct(query string) ([]domain.Product, error)
	CreateProduct(product utility.ProductRequest) (domain.Product, error)
}

type serviceProduct struct {
	repository RepositoryProduct
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
	case !utility.VerifyNonZeroValues(reqProduct):
		return domain.Product{}, utility.ErrInvalidValues
	case !utility.UniqueCodeValue(reqProduct.CodeValue, products):
		return domain.Product{}, utility.ErrUniqueCodeValue
	case !utility.VerifyExpirationDate(reqProduct.Expiration):
		return domain.Product{}, utility.ErrInvalidDate
	}

	return sp.repository.CreateProduct(reqProduct)
}

func NewServiceProduct(fileNameData string) *serviceProduct {

	rp := NewRepositoryProduct(fileNameData)

	return &serviceProduct{
		repository: rp,
	}
}

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
	UpdateProduct(pathVariable string, product utility.ProductRequest) (domain.Product, error)
	DeleteProduct(pathVariable string) error
	UpdatePatchProduct(pathVariable string, product utility.ProductPatchRequest) (domain.Product, error)
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

func NewServiceProduct(fileNameData string) *serviceProduct {

	rp := NewRepositoryProduct(fileNameData)

	return &serviceProduct{
		repository: rp,
	}
}

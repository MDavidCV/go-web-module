package product

import (
	"github.com/MDavidCV/go-web-module/internal/domain"
	"github.com/MDavidCV/go-web-module/utility"
)

type RepositoryProduct interface {
	GetProducts() ([]domain.Product, error)
	GetProductById(id int) (domain.Product, error)
	CreateProduct(product utility.ProductRequest) (domain.Product, error)
	UpdateProduct(int, utility.ProductRequest) (domain.Product, error)
	DeleteProduct(int) error
	UpdatePatchProduct(int, utility.ProductPatchRequest) (domain.Product, error)
}

type repositoryProduct struct {
	st map[int]domain.Product
}

func (rp *repositoryProduct) GetProducts() ([]domain.Product, error) {
	products := make([]domain.Product, 0, len(rp.st))
	for _, product := range rp.st {
		products = append(products, product)
	}

	if len(products) == 0 {
		return nil, utility.ErrNoProducts
	}

	return products, nil
}

func (rp *repositoryProduct) GetProductById(id int) (domain.Product, error) {
	product, ok := rp.st[id]

	if !ok {
		return domain.Product{}, utility.ErrProductNotFound
	}

	return product, nil
}

func (rp *repositoryProduct) CreateProduct(reqProduct utility.ProductRequest) (domain.Product, error) {
	id := len(rp.st) + 1
	product := domain.Product{
		Id:          id,
		Name:        reqProduct.Name,
		Quantity:    reqProduct.Quantity,
		CodeValue:   reqProduct.CodeValue,
		IsPublished: reqProduct.IsPublished,
		Expiration:  reqProduct.Expiration,
		Price:       reqProduct.Price,
	}

	if _, ok := rp.st[id]; ok {
		return domain.Product{}, utility.ErrProductAlreadyExists
	}

	rp.st[id] = product

	return product, nil
}

func (rp *repositoryProduct) UpdateProduct(id int, reqProduct utility.ProductRequest) (domain.Product, error) {
	product, ok := rp.st[id]

	if !ok {
		return domain.Product{}, utility.ErrProductNotFound
	}

	product.Name = reqProduct.Name
	product.Quantity = reqProduct.Quantity
	product.CodeValue = reqProduct.CodeValue
	product.IsPublished = reqProduct.IsPublished
	product.Expiration = reqProduct.Expiration
	product.Price = reqProduct.Price

	rp.st[id] = product

	return product, nil
}

func (rp *repositoryProduct) DeleteProduct(id int) error {
	if _, ok := rp.st[id]; !ok {
		return utility.ErrProductNotFound
	}

	delete(rp.st, id)
	return nil
}

func (rp *repositoryProduct) UpdatePatchProduct(id int, reqProduct utility.ProductPatchRequest) (domain.Product, error) {
	product, ok := rp.st[id]

	if !ok {
		return domain.Product{}, utility.ErrProductNotFound
	}

	if reqProduct.Name != nil {
		product.Name = *reqProduct.Name
	}

	if reqProduct.Quantity != nil {
		product.Quantity = *reqProduct.Quantity
	}

	if reqProduct.CodeValue != nil {
		product.CodeValue = *reqProduct.CodeValue
	}

	if reqProduct.IsPublished != nil {
		product.IsPublished = *reqProduct.IsPublished
	}

	if reqProduct.Expiration != nil {
		product.Expiration = *reqProduct.Expiration
	}

	if reqProduct.Price != nil {
		product.Price = *reqProduct.Price
	}

	rp.st[id] = product

	return product, nil
}

func NewRepositoryProduct(dataPath string) *repositoryProduct {
	st, err := utility.LoadProductsJson(dataPath)

	if err != nil {
		panic(err)
	}

	stMap := make(map[int]domain.Product, len(st))
	for _, product := range st {
		stMap[product.Id] = product
	}

	return &repositoryProduct{
		st: stMap,
	}
}

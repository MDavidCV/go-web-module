package product

import (
	"github.com/MDavidCV/go-web-module/internal/domain"
	"github.com/MDavidCV/go-web-module/utility"
)

type RepositoryProduct interface {
	GetProducts() ([]domain.Product, error)
	GetProductById(id int) (domain.Product, error)
	CreateProduct(product utility.ProductRequest) (domain.Product, error)
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

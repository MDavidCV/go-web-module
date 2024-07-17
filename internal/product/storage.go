package product

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/MDavidCV/go-web-module/internal/domain"
)

type StorageProduct interface {
	GetProducts() ([]domain.Product, error)
	WriteProducts(products map[int]domain.Product) error
}

type storageProduct struct {
	filename string
}

func (sp *storageProduct) GetProducts() ([]domain.Product, error) {
	f, err := os.Open(sp.filename)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
		return nil, err
	}
	defer f.Close()

	decoder := json.NewDecoder(f)

	var products []domain.Product

	// Read the open bracket
	if _, err := decoder.Token(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	// While the array contains values
	for decoder.More() {
		var product domain.Product
		// Decode one object
		if err := decoder.Decode(&product); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
			return nil, err
		}

		fmt.Println("Reading:", product)
		products = append(products, product)
	}

	// Read the closing bracket
	if _, err := decoder.Token(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return products, nil
}

func (sp *storageProduct) WriteProducts(productsMap map[int]domain.Product) error {

	products := make([]domain.Product, 0, len(productsMap))
	for _, product := range productsMap {
		products = append(products, product)
	}

	file, err := os.Create(sp.filename)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)

	if err := encoder.Encode(products); err != nil {
		log.Fatal(err)
		return err
	}

	//fmt.Println("JSON data written to file successfully.")
	return nil
}

func NewStorageProduct(filename string) *storageProduct {
	return &storageProduct{
		filename: filename,
	}
}

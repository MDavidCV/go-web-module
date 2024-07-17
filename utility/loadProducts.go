package utility

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/MDavidCV/go-web-module/internal/domain"
)

func LoadProductsJson(filename string) ([]domain.Product, error) {
	f, err := os.Open(filename)
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
	}

	// While the array contains values
	for decoder.More() {
		var product domain.Product
		// Decode one object
		if err := decoder.Decode(&product); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Reading:", product)
		products = append(products, product)
	}

	// Read the closing bracket
	if _, err := decoder.Token(); err != nil {
		log.Fatal(err)
	}

	return products, nil
}

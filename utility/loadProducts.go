package utility

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/MDavidCV/go-web-module/models"
)

func LoadProducts() []models.Product {
	f, err := os.Open("/Users/dcastrillonv/Documents/meli-boootcamp/go/go-web/go-web-module/products.json")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	defer f.Close()

	decoder := json.NewDecoder(f)

	var products []models.Product

	// Read the open bracket
	if _, err := decoder.Token(); err != nil {
		log.Fatal(err)
	}

	// While the array contains values
	for decoder.More() {
		var product models.Product
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

	return products
}

package utility

import (
	"time"

	"github.com/MDavidCV/go-web-module/dtos"
	"github.com/MDavidCV/go-web-module/models"
)

func VerifyNonZeroValues(body dtos.ProductRequest) bool {
	return body.Name != "" && body.Quantity != 0 && body.CodeValue != "" && body.Expiration != "" && body.Price != 0
}

func UniqueCodeValue(codeValue string, products map[int]models.Product) bool {
	for _, product := range products {
		if product.CodeValue == codeValue {
			return false
		}
	}
	return true
}

func VerifyExpirationDate(expiration string) bool {
	// Define the layout for parsing
	layout := "02/01/2006"
	_, err := time.Parse(layout, expiration)
	return err == nil
}

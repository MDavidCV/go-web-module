package utility

import (
	"time"

	"github.com/MDavidCV/go-web-module/internal/domain"
)

type ProductRequest struct {
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

func (pr *ProductRequest) VerifyNonZeroValues() bool {
	return pr.Name != "" && pr.Quantity != 0 && pr.CodeValue != "" && pr.Expiration != "" && pr.Price != 0
}

func (pr *ProductRequest) VerifyUniqueCodeValue(products []domain.Product) bool {
	for _, product := range products {
		if product.CodeValue == pr.CodeValue {
			return false
		}
	}
	return true
}

func (pr *ProductRequest) VerifyExpirationDate() bool {
	// Define the layout for parsing
	layout := "02/01/2006"
	_, err := time.Parse(layout, pr.Expiration)
	return err == nil
}

type ProductPatchRequest struct {
	Name        *string  `json:"name,omitempty"`
	Quantity    *int     `json:"quantity,omitempty"`
	CodeValue   *string  `json:"code_value,omitempty"`
	IsPublished *bool    `json:"is_published,omitempty"`
	Expiration  *string  `json:"expiration,omitempty"`
	Price       *float64 `json:"price,omitempty"`
}

func (ppr *ProductPatchRequest) VerifyNonZeroValues() bool {
	if ppr.Name != nil {
		return *ppr.Name != ""
	}

	if ppr.Quantity != nil {
		return *ppr.Quantity != 0
	}

	if ppr.CodeValue != nil {
		return *ppr.CodeValue != ""
	}

	if ppr.IsPublished != nil {
		return true
	}

	if ppr.Expiration != nil {
		return *ppr.Expiration != ""
	}

	if ppr.Price != nil {
		return *ppr.Price != 0
	}

	return false
}

func (ppr *ProductPatchRequest) VerifyUniqueCodeValue(products []domain.Product) bool {
	if ppr.CodeValue == nil {
		return true
	}

	for _, product := range products {
		if product.CodeValue == *ppr.CodeValue {
			return false
		}
	}
	return true
}

func (ppr *ProductPatchRequest) VerifyExpirationDate() bool {
	if ppr.Expiration == nil {
		return true
	}

	// Define the layout for parsing
	layout := "02/01/2006"
	_, err := time.Parse(layout, *ppr.Expiration)
	return err == nil
}

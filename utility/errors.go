package utility

import "errors"

var ErrProductNotFound = errors.New("product not found")
var ErrProductAlreadyExists = errors.New("product already exists")
var ErrInvalidQuery = errors.New("invalid query")
var ErrInvalidValues = errors.New("invalid values")
var ErrUniqueCodeValue = errors.New("code value already exists")
var ErrInvalidDate = errors.New("invalid expiration date")
var ErrInvalidId = errors.New("invalid id")
var ErrNoProducts = errors.New("no products found")
var ErrInvalidRequestBody = errors.New("invalid request body")

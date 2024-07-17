package utility

import (
	"net/http"
)

var errorCodes = map[error]int{
	ErrNoProducts:           http.StatusInternalServerError,
	ErrInvalidId:            http.StatusBadRequest,
	ErrProductNotFound:      http.StatusNotFound,
	ErrInvalidQuery:         http.StatusBadRequest,
	ErrInvalidDate:          http.StatusBadRequest,
	ErrUniqueCodeValue:      http.StatusBadRequest,
	ErrInvalidValues:        http.StatusBadRequest,
	ErrProductAlreadyExists: http.StatusInternalServerError,
}

type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"body"`
	Error string      `json:"error"`
}

func NewErrorResponse(err error) Response {
	return Response{
		Code:  errorCodes[err],
		Data:  nil,
		Error: err.Error(),
	}
}

func NewSuccessResponse(data interface{}) Response {
	return Response{
		Code:  http.StatusOK,
		Data:  data,
		Error: "",
	}
}

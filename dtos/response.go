package dtos

type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"body"`
	Error string      `json:"error"`
}

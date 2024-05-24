package responses

import (
	"encoding/json"
	"net/http"
)

type Object = map[string]interface{}

type SuccesfulResponse[Data any, Meta any] struct {
	Data    Data `json:"data"`
	Meta    Meta `json:"meta"`
	Success int  `json:"success"`
}

type ErrorfulResponse[Meta any] struct {
	Success int    `json:"success"`
	Error   string `json:"error"`
	Expired int    `json:"expired"`
	Meta    Meta   `json:"meta"`
}

func NewSuccessfulMetaResponse[Data any, Meta any](data Data, meta Meta) SuccesfulResponse[Data, Meta] {
	return SuccesfulResponse[Data, Meta]{
		Success: 1,
		Data:    data,
		Meta:    meta,
	}
}

func NewSuccessfulResponse[Data any](data Data) SuccesfulResponse[Data, []any] {
	return SuccesfulResponse[Data, []any]{
		Success: 1,
		Data:    data,
		Meta:    []any{},
	}
}

func (r SuccesfulResponse[Data, Meta]) Write(w http.ResponseWriter) {
	jsonData, err := json.Marshal(r)

	if err != nil {
		panic(err)
	}

	w.Write(jsonData)
}

func (r ErrorfulResponse[Meta]) Write(w http.ResponseWriter) {
	jsonData, err := json.Marshal(r)

	if err != nil {
		panic(err)
	}

	w.Write(jsonData)
}

func NewErrorfulMetaResponse[Meta any](message string, meta Meta) ErrorfulResponse[Meta] {
	return ErrorfulResponse[Meta]{
		Success: 0,
		Error:   message,
		Meta:    meta,
	}
}

func NewErrorfulResponse(message string) ErrorfulResponse[[]any] {
	return ErrorfulResponse[[]any]{
		Success: 0,
		Error:   message,
		Meta:    []any{},
	}
}

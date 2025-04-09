package handlers

import "net/http"

// This file contains a http.HandleFunc wrapper to always return a success or error.
// The "success" and "error" responses are defined in the "HandlerSuccess" and "HandlerError" structs
// and can be used as json responses.
// See indexHandler.go for an example
type ApiHandlerFunc func(w http.ResponseWriter, r *http.Request) (*HandlerSuccess, *HandlerError)

type HandlerSuccess struct {
	Status int `json:"-"`
	Data   interface{}
}

type HandlerError struct {
	Status  int `json:"-"`
	Message ErrorResponse
}

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

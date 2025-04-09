package handlers

import "net/http"

type IndexHandler struct {
}

func NewIndexHandler() *IndexHandler {
	return &IndexHandler{}
}

type healthResponse struct {
	Health string `json:"health"`
}

func (ih *IndexHandler) HealthCheck(w http.ResponseWriter, r *http.Request) (*HandlerSuccess, *HandlerError) {
	return &HandlerSuccess{Status: http.StatusOK, Data: healthResponse{Health: "Alive"}}, nil
}

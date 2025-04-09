package handlers

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
}

// User Response Model
type user struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

// Configuration of routes
func (uh *UserHandler) UserRouter() http.Handler {
	r := chi.NewRouter()

	// Middleware
	r.Use(logSomething)

	r.HandleFunc("GET /mock", ApiHandlerAdapter(uh.getMockUser))

	return r
}


func logSomething(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("This middleware should be applied only for user routes")
		next.ServeHTTP(w, r)
	})
}

func (uh *UserHandler) getMockUser(w http.ResponseWriter, r *http.Request) (*HandlerSuccess, *HandlerError) {
	shouldReturnUser := true

	if shouldReturnUser {
		return &HandlerSuccess{
			Status: http.StatusOK,
			Data:   &user{ID: "1", Name: "Yan", Email: "XO2iM@example.com"},
		}, nil
	}

	return nil, &HandlerError{
		Status:  http.StatusNotFound,
		Message: ErrorResponse{Code: "E404", Message: "User not found", Detail: ""},
	}
}

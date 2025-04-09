package server

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/hi-im-yan/crud-with-go/handlers"
)

type Server struct {
	Port   string
	Router *chi.Mux
}

func NewServer(port string) *Server {
	s := &Server{
		Port:   port,
		Router: chi.NewRouter(),
	}

	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.Recoverer)

	// Index Routes
	ih := handlers.NewIndexHandler()
	s.Router.HandleFunc("GET /", apiHandlerAdapter(ih.HealthCheck))

	return s
}

func (s *Server) Start() error {
	return http.ListenAndServe(":"+s.Port, s.Router)
}


// This function is a http.HandlerFunc adapter for my custom HandlerFunc called ApiHandlerFunc.
func apiHandlerAdapter(handler handlers.ApiHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		success, err := handler(w, r)

		if err != nil {
			w.WriteHeader(err.Status)
			json.NewEncoder(w).Encode(err.Message)
			return
		}

		if success != nil {
			w.WriteHeader(success.Status)
			json.NewEncoder(w).Encode(success.Data)
		}
	}
}

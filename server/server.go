package server

import (
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
	s.Router.HandleFunc("GET /", handlers.ApiHandlerAdapter(ih.HealthCheck))

	// User Routes
	uh := handlers.NewUserHandler()
	s.Router.Mount("/users", uh.UserRouter())

	return s
}

func (s *Server) Start() error {
	return http.ListenAndServe(":"+s.Port, s.Router)
}

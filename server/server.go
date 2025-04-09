package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/hi-im-yan/crud-with-go/handlers"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Server struct {
	Port   string
	Router *chi.Mux
	DB     *pgxpool.Pool
}

func NewServer(port string, db *pgxpool.Pool) *Server {
	s := &Server{
		Port:   port,
		Router: chi.NewRouter(),
		DB:     db,
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

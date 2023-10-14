package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	router *chi.Mux
	db     *sqlx.DB
}

func NewServer(dbConn *sqlx.DB) *Server {
	return &Server{db: dbConn, router: chi.NewRouter()}
}

func (server *Server) Run() error {
	return http.ListenAndServe(":8080", server.router)
}

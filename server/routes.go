package server

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
)

func ConfigureRoutes(server *Server) {
	server.router.Use(middleware.Logger)
	server.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Heeeeelloo :)!!"))
	})
}

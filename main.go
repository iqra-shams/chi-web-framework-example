package main

import (
	"net/http"

	"github.com/iqra-shams/chi/api"
	"github.com/iqra-shams/chi/handler"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Index handler
	r.Post("/", api.ProtectedHandler)
	r.Post("/login",handler.LoginHandler)

	http.ListenAndServe(":3333", r)
}

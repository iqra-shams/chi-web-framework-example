package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/iqra-shams/chi/handler"
)

func main() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Index handler

	r.Post("/res", handler.Restricted)
	r.Post("/login", handler.LoginHandler)

	http.ListenAndServe(":3333", r)
}

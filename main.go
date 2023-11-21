package main

import (
	"github.com/iqra-shams/chi/api"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Index handler
	r.Post("/", api.HandlerPostReq)

	http.ListenAndServe(":3333", r)

}

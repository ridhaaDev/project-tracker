package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func helloController(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome"))
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", helloController)
	http.ListenAndServe(":3000", r)
}

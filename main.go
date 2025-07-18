package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"terraform_http_backend/src"
	"terraform_http_backend/src/handler"
	"terraform_http_backend/src/log"
)

func main() {
	log.Init()
	log.Info("Starting %s", src.AppName)
	r := chi.NewRouter()
	r.Get("/state/{projectName}/{environment}", handler.GetState)
	r.Post("/state/{projectName}/{environment}", handler.SetState)

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}

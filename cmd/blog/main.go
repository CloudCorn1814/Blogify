package main

import (
	"Blogify/db"
	"Blogify/internal/blog/handler"
	"Blogify/internal/blog/repository"
	"Blogify/internal/blog/service"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	db, err := db.InitDB("user", "password", "db")
	if err != nil {
		log.Fatalf("Could not set up database: %v", err)
	}

	repo := repository.NewRepository(db)
	srvs := service.NewService(repo)
	handler := handler.NewHandler(srvs)

	r := chi.NewRouter()
	r.Post("/article", handler.HandleCreate)
	r.Get("/article", handler.HandleGetAll)
	r.Put("/article/{articleID}", handler.HandleUpdate)
	r.Get("/article/{articleID}", handler.HandleGet)
	r.Delete("/article/{articleID}", handler.HandleDelete)

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("server start error")
	}
}

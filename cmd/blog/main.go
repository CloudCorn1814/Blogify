package main

import (
	"Blogify/db"
	"Blogify/internal/blog/handler"
	"Blogify/internal/blog/repository"
	"Blogify/internal/blog/service"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PostgresPort := os.Getenv("POSTGRES_PORT")
	PostgresHost := os.Getenv("POSTGRES_HOST")
	LocalPort := os.Getenv("PORT")
	Username := os.Getenv("USERNAME")
	Password := os.Getenv("PASSWORD")
	DB := os.Getenv("DB")

	db, err := db.InitDB(PostgresPort, PostgresHost, Username, Password, DB)
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

	err = http.ListenAndServe(LocalPort, r)
	if err != nil {
		log.Fatal("server start error")
	}
}

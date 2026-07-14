package main

import (
	"Blogify/contracts/gen/Blogify/contracts/gen"
	"Blogify/db"
	"Blogify/internal/blog/handler"
	"Blogify/internal/blog/middleware"
	"Blogify/internal/blog/repository"
	"Blogify/internal/blog/service"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

	conn, err := grpc.NewClient(os.Getenv("gRPC_server"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("invalid grpc conn: %v", err)
	}
	client := gen.NewAuthClient(conn)
	middleware := middleware.NewMiddleware(client)
	r := chi.NewRouter()

	r.Route("/article", func(r chi.Router) {
		r.Get("/", handler.HandleGetAll)
		r.Get("/{articleID}", handler.HandleGet)
	})

	r.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleware)
		r.Route("/article", func(r chi.Router) {
			r.Post("/", handler.HandleCreate)
			r.Put("/{articleID}", handler.HandleUpdate)
			r.Delete("/{articleID}", handler.HandleDelete)
		})
	})

	err = http.ListenAndServe(LocalPort, r)
	if err != nil {
		log.Fatal("server start error")
	}
}

package main

import (
	"Blogify/contracts/gen/Blogify/contracts/gen"
	"Blogify/db"
	"Blogify/internal/auth/handler"
	"Blogify/internal/auth/repository"
	"Blogify/internal/auth/server"
	"Blogify/internal/auth/service"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
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
	grpcServer := grpc.NewServer()
	gen.RegisterAuthServer(grpcServer, server.NewServer(srvs))

	r := chi.NewRouter()
	r.Post("/sign_in", handler.HandleCreate)
	r.Post("/login", handler.HandleLogin)

	go func() {
		lis, err := net.Listen("tcp", ":9090")
		if err != nil {
			log.Fatal(err)
		}
		grpcServer.Serve(lis)
	}()

	err = http.ListenAndServe(LocalPort, r)
	if err != nil {
		log.Fatal("server start error")
	}
}

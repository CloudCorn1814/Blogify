package main

import (
	"Blogify/contracts/gen/Blogify/contracts/gen"
	"Blogify/db"
	"Blogify/internal/auth/handler"
	"Blogify/internal/auth/producer"
	"Blogify/internal/auth/repository"
	"Blogify/internal/auth/server"
	"Blogify/internal/auth/service"
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	db, err := db.InitDB(PostgresHost, PostgresPort, Username, Password, DB)
	if err != nil {
		log.Fatalf("Could not set up database: %v", err)
	}

	repo := repository.NewRepository(db)
	producer, err := producer.NewProducer(os.Getenv("KAFKA"))
	if err != nil {
		log.Fatalf("producer init error: %v", err)
	}
	defer producer.Close()

	srvs := service.NewService(repo, producer)
	handler := handler.NewHandler(srvs)
	grpcServer := grpc.NewServer()
	gen.RegisterAuthServer(grpcServer, server.NewServer(srvs))

	r := chi.NewRouter()
	r.Route("/api/v1/auth", func(r chi.Router) {
		r.Post("/register", handler.HandleCreate)
		r.Post("/login", handler.HandleLogin)
	})

	go func() {
		lis, err := net.Listen("tcp", ":9090")
		if err != nil {
			log.Fatal(err)
		}
		grpcServer.Serve(lis)
	}()
	srv := &http.Server{Addr: LocalPort, Handler: r}
	go func() {
		err = srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal("server start error")
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	log.Println("Shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	grpcServer.GracefulStop()

	err = srv.Shutdown(ctx)
	if err != nil {
		log.Fatal("server shutdown error")
	}
}

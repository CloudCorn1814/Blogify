package main

import (
	"Blogify/contracts/gen/Blogify/contracts/gen"
	"Blogify/db"
	"Blogify/internal/blog/handler"
	"Blogify/internal/blog/middleware"
	"Blogify/internal/blog/repository"
	"Blogify/internal/blog/service"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	db, err := db.InitDB(PostgresHost, PostgresPort, Username, Password, DB)
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
	err = srv.Shutdown(ctx)
	if err != nil {
		log.Fatal("server shutdown error")
	}
}

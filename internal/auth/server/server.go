package server

import (
	"Blogify/contracts/gen/Blogify/contracts/gen"
	"Blogify/internal/auth/service"
	"context"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type Server struct {
	service *service.Service
	gen.UnimplementedAuthServer
}

func NewServer(service *service.Service) *Server {
	return &Server{service: service}
}

func (s *Server) Check(ctx context.Context, req *gen.CheckRequest) (*gen.CheckResponse, error) {
	token, err := jwt.Parse(req.Token, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("KEY")), nil
	})
	if err != nil {
		return nil, fmt.Errorf("token parsing error: %w", err)
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid claims")
	}
	userID := claims["id"].(int64)
	return &gen.CheckResponse{Id: userID}, nil
}

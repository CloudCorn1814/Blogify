package middleware

import (
	"Blogify/contracts/gen/Blogify/contracts/gen"
	"context"
	"net/http"
)

type Middleware struct {
	c gen.AuthClient
}

func NewMiddleware(c gen.AuthClient) *Middleware {
	return &Middleware{c: c}
}

func (m *Middleware) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "empty token", http.StatusBadRequest)
			return
		}
		userID, err := m.c.Check(r.Context(), &gen.CheckRequest{Token: authHeader})
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), "userID", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

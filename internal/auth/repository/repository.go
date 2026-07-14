package repository

import (
	"Blogify/db"
	"Blogify/internal/auth/entity"
)

type Repository struct {
	db db.Database
}

func NewRepository(db db.Database) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateUser(login, passwordHash string) (*entity.User, error) {
	var id int
	query := `INSERT INTO users (login, password_hash)  VALUES ($1, $2) RETURNING id`
	err := r.db.Conn.QueryRow(query, login, passwordHash).Scan(&id)
	if err != nil {
		return nil, err
	}
	return &entity.User{ID: id, Login: login, PasswordHash: passwordHash}, nil
}

func (r *Repository) LoginUser(login string) (*int, string, error) {
	var passwordHash string
	var id int
	query := `SELECT id, password_hash FROM users WHERE login=$1`
	err := r.db.Conn.QueryRow(query, login).Scan(&id, &passwordHash)
	if err != nil {
		return nil, "", err
	}
	return &id, passwordHash, nil
}

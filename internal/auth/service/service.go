package service

import (
	"Blogify/internal/auth/entity"
	"Blogify/internal/auth/producer"
	"Blogify/internal/auth/repository"
	"errors"
	"fmt"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo     *repository.Repository
	producer *producer.Producer
}

func NewService(repo *repository.Repository, producer *producer.Producer) *Service {
	return &Service{repo: repo, producer: producer}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (s *Service) CreateUser(login, password string) (*entity.User, error) {
	pwd, err := HashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("hashing error: %w", err)
	}
	user, err := s.repo.CreateUser(login, pwd)
	if err != nil {
		return nil, fmt.Errorf("user creation error: %w", err)
	}
	err = s.producer.SendMessage("user.registered", login)
	if err != nil {
		return nil, fmt.Errorf("something went wrong: %w", err)
	}
	return user, nil
}

func (s *Service) LoginUser(login, password string) (string, error) {
	id, hash, err := s.repo.LoginUser(login)
	if err != nil {
		return "", fmt.Errorf("login error: %w", err)
	}
	expire := time.Now().Add(24 * time.Hour).Unix()
	if !CheckPasswordHash(password, hash) {
		return "", errors.New("password is invalid")
	}
	key := os.Getenv("KEY")
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":  id,
			"exp": expire,
		},
	)
	signet, err := t.SignedString([]byte(key))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}
	return signet, nil
}

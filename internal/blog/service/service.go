package service

import (
	"Blogify/internal/blog/entity"
	"Blogify/internal/blog/repository"
	"errors"
)

var (
	errorEmptyField = errors.New("field empty")
)

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateArticle(author, topic, text string) (*entity.Article, error) {
	if author == "" || topic == "" || text == "" {
		return nil, errorEmptyField
	}
	article, err := s.repo.CreateArticle(author, topic, text)
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (s *Service) UpdateArticle(article entity.Article) (*entity.Article, error) {
	articleNew, err := s.repo.UpdateArticle(article)
	if err != nil {
		return nil, err
	}
	return articleNew, nil
}

func (s *Service) GetArticle(ID int) (*entity.Article, error) {
	article, err := s.repo.GetArticle(ID)
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (s *Service) GetArticleAll() ([]entity.Article, error) {
	articles, err := s.repo.GetArticleAll()
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (s *Service) DeleteArticle(ID int) error {
	err := s.repo.DeleteArticle(ID)
	if err != nil {
		return err
	}
	return nil
}

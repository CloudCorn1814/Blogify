package service

import (
	"Blogify/internal/blog/entity"
	"Blogify/internal/blog/repository"
	"fmt"
)

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateArticle(author, topic, text string) (*entity.Article, error) {
	article, err := s.repo.CreateArticle(author, topic, text)
	if err != nil {
		return nil, fmt.Errorf("article creation error: %w", err)
	}
	return article, nil
}

func (s *Service) UpdateArticle(article entity.Article) (*entity.Article, error) {
	articleNew, err := s.repo.UpdateArticle(article)
	if err != nil {
		return nil, fmt.Errorf("article update error: %w", err)
	}
	return articleNew, nil
}

func (s *Service) GetArticle(ID int) (*entity.Article, error) {
	article, err := s.repo.GetArticle(ID)
	if err != nil {
		return nil, fmt.Errorf("article fetch error: %w", err)
	}
	return article, nil
}

func (s *Service) GetArticleAll() ([]entity.Article, error) {
	articles, err := s.repo.GetArticleAll()
	if err != nil {
		return nil, fmt.Errorf("articles fetch error: %w", err)
	}
	return articles, nil
}

func (s *Service) DeleteArticle(ID int) error {
	err := s.repo.DeleteArticle(ID)
	if err != nil {
		return fmt.Errorf("deletion error: %w", err)
	}
	return nil
}

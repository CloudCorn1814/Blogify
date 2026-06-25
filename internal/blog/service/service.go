package service

import (
	"Blogify/internal/blog/entity"
	"Blogify/internal/blog/repository"
)

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateArticle(author, topic, text string) (article entity.Article, err error) {
	return article, nil
}

func (s *Service) UpdateArticle(article entity.Article) (articleNew entity.Article, err error) {
	return articleNew, nil
}

func (s *Service) GetArticle(ID int) (article entity.Article, err error) {
	return article, nil
}

func (s *Service) GetArticleAll() (articles []entity.Article, err error) {
	return articles, nil
}

func (r *Service) DeleteArticle(ID int) error {
	return nil
}

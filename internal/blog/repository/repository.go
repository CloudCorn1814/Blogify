package repository

import (
	"Blogify/db"
	"Blogify/internal/blog/entity"
)

type Repository struct {
	db db.Database
}

func NewRepository(db db.Database) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateArticle(author, topic, text string) (article entity.Article, err error) {
	return article, nil
}

func (r *Repository) UpdateArticle(article entity.Article) (articleNew entity.Article, err error) {
	return articleNew, nil
}

func (r *Repository) GetArticle(ID int) (article entity.Article, err error) {
	return article, nil
}

func (r *Repository) GetArticleAll() (articles []entity.Article, err error) {
	return articles, nil
}

func (r *Repository) DeleteArticle(ID int) error {
	return nil
}

package repository

import (
	"Blogify/db"
	"Blogify/internal/blog/entity"
	"time"
)

type Repository struct {
	db db.Database
}

func NewRepository(db db.Database) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateArticle(author, topic, text string) (*entity.Article, error) {
	var id int
	var createdAt time.Time
	query := `INSERT INTO articles (author, topic, text) VALUES ($1, $2, $3) RETURNING id, created_at`
	err := r.db.Conn.QueryRow(query, author, topic, text).Scan(&id, &createdAt)
	if err != nil {
		return nil, err
	}
	return &entity.Article{ID: id, Author: author, Topic: topic, Text: text, PostedTime: createdAt}, nil
}

func (r *Repository) UpdateArticle(article entity.Article) (*entity.Article, error) {
	var changedAt time.Time
	query := `UPDATE articles SET author=$1, topic=$2, text=$3 WHERE id=$4 RETURNING changed_at`
	err := r.db.Conn.QueryRow(query, article.Author, article.Topic, article.Text, article.ID).Scan(&changedAt)
	if err != nil {
		return nil, err
	}
	return &entity.Article{ID: article.ID, Author: article.Author, Topic: article.Topic, Text: article.Text, PostedTime: changedAt}, nil
}

func (r *Repository) GetArticle(ID int) (*entity.Article, error) {
	var article entity.Article
	query := `SELECT * FROM articles WHERE id=$1`
	err := r.db.Conn.QueryRow(query, ID).Scan(&article.ID, &article.Author, &article.Topic, &article.Text, &article.PostedTime)
	if err != nil {
		return nil, err
	}
	return &article, nil
}

func (r *Repository) GetArticleAll() ([]entity.Article, error) {
	query := `SELECT * FROM articles`
	rows, err := r.db.Conn.Query(query)
	if err != nil {
		return nil, err
	}
	var articles []entity.Article
	for rows.Next() {
		var article entity.Article
		err := rows.Scan(&article.ID, &article.Author, &article.Topic, &article.Text, &article.PostedTime)
		if err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}
	return articles, nil
}

func (r *Repository) DeleteArticle(ID int) error {
	query := `DELETE FROM articles WHERE id=$1`
	_, err := r.db.Conn.Exec(query, ID)
	if err != nil {
		return err
	}
	return nil
}

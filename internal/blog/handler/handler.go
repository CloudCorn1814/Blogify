package handler

import (
	"Blogify/internal/blog/entity"
	"Blogify/internal/blog/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func urlID(w http.ResponseWriter, r *http.Request, url string) (int, error) {
	id, err := strconv.Atoi(chi.URLParam(r, url))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 0, fmt.Errorf("ID fetch error: %w", err)
	}
	return id, nil
}

func (h *Handler) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var articleRequest createArticleRequest
	err := json.NewDecoder(r.Body).Decode(&articleRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	article, err := h.service.CreateArticle(articleRequest.Author, articleRequest.Text, articleRequest.Topic)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(article)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	var articleRequest updateArticleArticleRequest
	err := json.NewDecoder(r.Body).Decode(&articleRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	articleID, err := urlID(w, r, "articleID")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	articleBuild := entity.Article{
		ID:     articleID,
		Author: articleRequest.Author,
		Topic:  articleRequest.Topic,
		Text:   articleRequest.Text,
	}
	article, err := h.service.UpdateArticle(articleBuild)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(article)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) HandleGet(w http.ResponseWriter, r *http.Request) {
	articleID, err := urlID(w, r, "articleID")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	article, err := h.service.GetArticle(articleID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(article)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) HandleGetAll(w http.ResponseWriter, r *http.Request) {
	articles, err := h.service.GetArticleAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(articles)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) HandleDelete(w http.ResponseWriter, r *http.Request) {
	articleID, err := urlID(w, r, "articleID")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.service.DeleteArticle(articleID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode("article deleted")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

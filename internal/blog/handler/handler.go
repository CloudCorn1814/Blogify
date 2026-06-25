package handler

import (
	"Blogify/internal/blog/service"
	"net/http"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) HandleCreate(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) HandleUpdate(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) HandleGet(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) HandleGetAll(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) HandleDelete(w http.ResponseWriter, r *http.Request) {}

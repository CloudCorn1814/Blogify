package handler

import (
	"Blogify/internal/auth/service"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Handler struct {
	service  *service.Service
	validate *validator.Validate
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service, validate: validator.New(validator.WithRequiredStructEnabled())}
}

func (h *Handler) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var reqCredits CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&reqCredits)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.validate.Struct(reqCredits)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newUser, err := h.service.CreateUser(reqCredits.Login, reqCredits.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var reqCredits LoginUserRequest
	err := json.NewDecoder(r.Body).Decode(&reqCredits)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.validate.Struct(reqCredits)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	signet, err := h.service.LoginUser(reqCredits.Login, reqCredits.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(signet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

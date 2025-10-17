package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gplaydb/internal/models"
	"gplaydb/internal/services"

	"github.com/google/uuid"
)

type UserHandler struct {
	Service *services.UserService
}

func NewUserHandler(s *services.UserService) *UserHandler {
	return &UserHandler{Service: s}
}

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.Service.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	userUUID, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Erro ao passar o id de string para UUID", http.StatusBadRequest)
		return
	}

	user, err := h.Service.GetUserById(userUUID)
	if err != nil {
		http.Error(w, "Usuário não encontrado", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)

}

func (h *UserHandler) InsertUser(w http.ResponseWriter, r *http.Request) {

	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	u, err := h.Service.InsertUser(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)

}

func (h *UserHandler) DeleteUserById(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	userUUID, err1 := uuid.Parse(idStr)
	if err1 != nil {
		http.Error(w, "Erro ao passar o id de string para UUID", http.StatusBadRequest)
		return
	}

	err2 := h.Service.DeleteUserById(userUUID)
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	userUUID, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Erro ao passar o id de string para UUID", http.StatusBadRequest)
		return
	}

	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user.ID = userUUID
	u, err := h.Service.UpdateUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(u)

}

func (h *UserHandler) UserWithProducts(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	userUUID, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Erro ao passar o id de string para UUID", http.StatusBadRequest)
		return
	}

	user, err := h.Service.UserWithProducts(userUUID)
	fmt.Print(err)
	if err != nil {
		http.Error(w, "Usuário não encontrado", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

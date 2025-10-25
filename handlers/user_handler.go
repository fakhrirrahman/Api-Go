package handlers

import (
	"encoding/json"
	"goApi/models"
	"goApi/response"
	"goApi/services"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	services *services.UserService
}

func NewUserHandler(s *services.UserService) *UserHandler {
	return &UserHandler{services: s}
}

// get user
func (h *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.services.ListUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response.JSON(w, http.StatusOK, "Berhasil Mendapatkan semua user", users)
}

// get user by id
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, "ID user tidak valid", nil)
		return
	}

	user, err := h.services.GetUserrByID(id)
	if err != nil {
		response.JSON(w, http.StatusNotFound, "User tidak ditemukan", nil)
		return
	}

	response.JSON(w, http.StatusOK, "Berhasil mendapatkan data user", user)
}

// create user
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var u models.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "Gagal decode data user", http.StatusBadRequest)
		return
	}
	if err := h.services.AddUser(u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response.JSON(w, http.StatusCreated, "User berhasil dibuat", u)
}

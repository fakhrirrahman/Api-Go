package handlers

import (
	"encoding/json"
	"goApi/models"
	"goApi/response"
	"goApi/services"
	"goApi/utils"
	"net/http"
)

type AuthHandler struct {
	userService *services.UserService
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token  string      `json:"token"`
	User   models.User `json:"user"`
	Expiry string      `json:"expiry"`
}

func NewAuthHandler(userService *services.UserService) *AuthHandler {
	return &AuthHandler{userService: userService}
}

// Login endpoint - dapatkan JWT token
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginReq LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		response.JSON(w, http.StatusBadRequest, "Format request tidak valid", nil)
		return
	}

	// Validasi input
	if loginReq.Email == "" || loginReq.Password == "" {
		response.JSON(w, http.StatusBadRequest, "Email dan password diperlukan", nil)
		return
	}

	// Di sini Anda bisa validate user terhadap database
	// Untuk demo, kami accept semua request dengan format valid
	// Dalam production, cek password terhadap hash di database

	// Generate JWT token
	token, err := utils.GenerateToken(1, loginReq.Email, "User")
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, "Gagal generate token", nil)
		return
	}

	loginRes := LoginResponse{
		Token: token,
		User: models.User{
			ID:    1,
			Email: loginReq.Email,
			Name:  "User",
		},
		Expiry: "24 jam",
	}

	response.JSON(w, http.StatusOK, "Login berhasil", loginRes)
}

// Register endpoint
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		response.JSON(w, http.StatusBadRequest, "Format request tidak valid", nil)
		return
	}

	if err := h.userService.AddUser(user); err != nil {
		response.JSON(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	// Generate token untuk user baru
	token, err := utils.GenerateToken(user.ID, user.Email, user.Name)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, "Gagal generate token", nil)
		return
	}

	registerRes := LoginResponse{
		Token:  token,
		User:   user,
		Expiry: "24 jam",
	}

	response.JSON(w, http.StatusCreated, "Register berhasil", registerRes)
}

package handlers

import (
	"goApi/models"
	"goApi/services"
	"goApi/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
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

// LoginFiber godoc
// @Summary Login user
// @Description Login with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param login body LoginRequest true "Login credentials"
// @Success 200 {object} LoginResponse
// @Router /auth/login [post]
// LoginFiber - Fiber version
func (h *AuthHandler) LoginFiber(c *fiber.Ctx) error {
	var loginReq LoginRequest

	if err := c.BodyParser(&loginReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Format request tidak valid",
			"data":    nil,
		})
	}

	if loginReq.Email == "" || loginReq.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Email dan password diperlukan",
			"data":    nil,
		})
	}

	token, err := utils.GenerateToken(1, loginReq.Email, "User")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Gagal generate token",
			"data":    nil,
		})
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

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Login berhasil",
		"data":    loginRes,
	})
}

// RegisterFiber godoc
// @Summary Register user
// @Description Register a new user
// @Tags auth
// @Accept json
// @Produce json
// @Param user body models.User true "User data"
// @Success 201 {object} LoginResponse
// @Router /auth/register [post]
// RegisterFiber - Fiber version
func (h *AuthHandler) RegisterFiber(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Format request tidak valid",
			"data":    nil,
		})
	}

	if err := h.userService.AddUser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"data":    nil,
		})
	}

	token, err := utils.GenerateToken(user.ID, user.Email, user.Name)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Gagal generate token",
			"data":    nil,
		})
	}

	registerRes := LoginResponse{
		Token:  token,
		User:   user,
		Expiry: "24 jam",
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  fiber.StatusCreated,
		"message": "Register berhasil",
		"data":    registerRes,
	})
}

// Legacy methods untuk compatibility
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	// Deprecated: Use LoginFiber instead
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	// Deprecated: Use RegisterFiber instead
}

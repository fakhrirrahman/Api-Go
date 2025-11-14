package handlers

import (
	"goApi/models"
	"goApi/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	services *services.UserService
}

func NewUserHandler(s *services.UserService) *UserHandler {
	return &UserHandler{services: s}
}

// GetUserFiber - Fiber version
func (h *UserHandler) GetUserFiber(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "ID user tidak valid",
			"data":    nil,
		})
	}

	user, err := h.services.GetUserrByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": "User tidak ditemukan",
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Berhasil mendapatkan data user",
		"data":    user,
	})
}

// CreateUserFiber - Fiber version
func (h *UserHandler) CreateUserFiber(c *fiber.Ctx) error {
	var u models.User
	if err := c.BodyParser(&u); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Gagal decode data user",
			"data":    nil,
		})
	}

	if err := h.services.AddUser(u); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  fiber.StatusCreated,
		"message": "User berhasil dibuat",
		"data":    u,
	})
}

// ListUsersFiber - Fiber version
func (h *UserHandler) ListUsersFiber(c *fiber.Ctx) error {
	users, err := h.services.ListUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Berhasil Mendapatkan semua user",
		"data":    users,
	})
}

// Legacy methods untuk compatibility
func (h *UserHandler) ListUsers(w interface{}, r interface{})  {}
func (h *UserHandler) GetUser(w interface{}, r interface{})    {}
func (h *UserHandler) CreateUser(w interface{}, r interface{}) {}

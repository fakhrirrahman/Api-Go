package handlers

import (
	"goApi/models"
	"goApi/services"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type BookingHandler struct {
	service *services.BookingService
}

func NewBookingHandler(s *services.BookingService) *BookingHandler {
	return &BookingHandler{service: s}
}

// CreateBookingFiber handles POST /bookings
func (h *BookingHandler) CreateBookingFiber(c *fiber.Ctx) error {
	var req models.CreateBookingRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Format data booking tidak valid",
			"data":    nil,
		})
	}

	id, err := h.service.CreateBooking(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"data":    nil,
		})
	}

	booking, _ := h.service.GetBooking(id)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  fiber.StatusCreated,
		"message": "Booking berhasil dibuat",
		"data":    booking,
	})
}

// GetBookingFiber handles GET /bookings/:id
func (h *BookingHandler) GetBookingFiber(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "ID booking tidak valid",
			"data":    nil,
		})
	}

	booking, err := h.service.GetBooking(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": "Booking tidak ditemukan",
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Berhasil mendapatkan booking",
		"data":    booking,
	})
}

// ListBookingsFiber handles GET /bookings
func (h *BookingHandler) ListBookingsFiber(c *fiber.Ctx) error {
	bookings, err := h.service.ListBookings()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Gagal mendapatkan list booking",
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Berhasil mendapatkan list booking",
		"data":    bookings,
	})
}

// CheckAvailabilityFiber handles GET /bookings/check-availability?field_id=1&start_time=...&end_time=...
func (h *BookingHandler) CheckAvailabilityFiber(c *fiber.Ctx) error {
	fieldIDStr := c.Query("field_id")
	startTimeStr := c.Query("start_time")
	endTimeStr := c.Query("end_time")

	if fieldIDStr == "" || startTimeStr == "" || endTimeStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Parameter field_id, start_time, dan end_time diperlukan",
			"data":    nil,
		})
	}

	fieldID, err := strconv.Atoi(fieldIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Field ID tidak valid",
			"data":    nil,
		})
	}

	startTime, err := parseTime(startTimeStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Format start_time tidak valid",
			"data":    nil,
		})
	}

	endTime, err := parseTime(endTimeStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Format end_time tidak valid",
			"data":    nil,
		})
	}

	available, err := h.service.CheckAvailability(fieldID, startTime, endTime)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Pengecekan ketersediaan berhasil",
		"data": fiber.Map{
			"available": available,
		},
	})
}

func parseTime(timeStr string) (time.Time, error) {
	return time.Parse(time.RFC3339, timeStr)
}

// Legacy methods untuk compatibility
func (h *BookingHandler) CreateBooking(w interface{}, r interface{}) {}
func (h *BookingHandler) GetBooking(w interface{}, r interface{})    {}
func (h *BookingHandler) ListBookings(w interface{}, r interface{}) {}
func (h *BookingHandler) CheckAvailability(w interface{}, r interface{}) {}

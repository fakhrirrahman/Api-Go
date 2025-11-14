package handlers

import (
	"goApi/models"
	"goApi/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type PaymentHandler struct {
	service *services.PaymentService
}

func NewPaymentHandler(s *services.PaymentService) *PaymentHandler {
	return &PaymentHandler{service: s}
}

// ProcessPaymentFiber godoc
// @Summary Process payment
// @Description Process payment for a booking
// @Tags payments
// @Accept json
// @Produce json
// @Param payment body models.ProcessPaymentRequest true "Payment request"
// @Success 200 {object} models.Payment
// @Router /payments [post]
// ProcessPaymentFiber handles POST /payments
func (h *PaymentHandler) ProcessPaymentFiber(c *fiber.Ctx) error {
	var req models.ProcessPaymentRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Format data payment tidak valid",
			"data":    nil,
		})
	}

	if req.BookingID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "booking_id diperlukan",
			"data":    nil,
		})
	}

	payment, err := h.service.ProcessPayment(req.BookingID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Payment berhasil diproses",
		"data":    payment,
	})
}

// GetPaymentFiber godoc
// @Summary Get payment by ID
// @Description Get a specific payment
// @Tags payments
// @Produce json
// @Param id path int true "Payment ID"
// @Success 200 {object} models.Payment
// @Router /payments/{id} [get]
// GetPaymentFiber handles GET /payments/:id
func (h *PaymentHandler) GetPaymentFiber(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "ID payment tidak valid",
			"data":    nil,
		})
	}

	payment, err := h.service.GetPayment(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": "Payment tidak ditemukan",
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Berhasil mendapatkan data payment",
		"data":    payment,
	})
}

// GetPaymentByBookingFiber godoc
// @Summary Get payment by booking ID
// @Description Get payment for a specific booking
// @Tags payments
// @Produce json
// @Param booking_id path int true "Booking ID"
// @Success 200 {object} models.Payment
// @Router /payments/booking/{booking_id} [get]
// GetPaymentByBookingFiber handles GET /payments/booking/:booking_id
func (h *PaymentHandler) GetPaymentByBookingFiber(c *fiber.Ctx) error {
	bookingIDStr := c.Params("booking_id")
	bookingID, err := strconv.Atoi(bookingIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Booking ID tidak valid",
			"data":    nil,
		})
	}

	payment, err := h.service.GetPaymentByBooking(bookingID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": "Payment untuk booking tidak ditemukan",
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Berhasil mendapatkan payment",
		"data":    payment,
	})
}

// Legacy methods untuk compatibility
func (h *PaymentHandler) ProcessPayment(w interface{}, r interface{})      {}
func (h *PaymentHandler) GetPayment(w interface{}, r interface{})          {}
func (h *PaymentHandler) GetPaymentByBooking(w interface{}, r interface{}) {}

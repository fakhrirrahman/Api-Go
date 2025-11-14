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

type PaymentHandler struct {
	service *services.PaymentService
}

func NewPaymentHandler(s *services.PaymentService) *PaymentHandler {
	return &PaymentHandler{service: s}
}

// ProcessPayment handles POST /payments
func (h *PaymentHandler) ProcessPayment(w http.ResponseWriter, r *http.Request) {
	var req models.ProcessPaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.JSON(w, http.StatusBadRequest, "Format data payment tidak valid", nil)
		return
	}

	if req.BookingID == 0 {
		response.JSON(w, http.StatusBadRequest, "booking_id diperlukan", nil)
		return
	}

	payment, err := h.service.ProcessPayment(req.BookingID)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	response.JSON(w, http.StatusOK, "Payment berhasil diproses", payment)
}

// GetPayment handles GET /payments/{id}
func (h *PaymentHandler) GetPayment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, "ID payment tidak valid", nil)
		return
	}

	payment, err := h.service.GetPayment(id)
	if err != nil {
		response.JSON(w, http.StatusNotFound, "Payment tidak ditemukan", nil)
		return
	}

	response.JSON(w, http.StatusOK, "Berhasil mendapatkan data payment", payment)
}

// GetPaymentByBooking handles GET /payments/booking/{booking_id}
func (h *PaymentHandler) GetPaymentByBooking(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookingIDStr := params["booking_id"]
	bookingID, err := strconv.Atoi(bookingIDStr)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, "Booking ID tidak valid", nil)
		return
	}

	payment, err := h.service.GetPaymentByBooking(bookingID)
	if err != nil {
		response.JSON(w, http.StatusNotFound, "Payment untuk booking tidak ditemukan", nil)
		return
	}

	response.JSON(w, http.StatusOK, "Berhasil mendapatkan payment", payment)
}

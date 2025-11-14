package handlers

import (
	"encoding/json"
	"goApi/models"
	"goApi/response"
	"goApi/services"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type BookingHandler struct {
	service *services.BookingService
}

func NewBookingHandler(s *services.BookingService) *BookingHandler {
	return &BookingHandler{service: s}
}

// CreateBooking handles POST /bookings
func (h *BookingHandler) CreateBooking(w http.ResponseWriter, r *http.Request) {
	var req models.CreateBookingRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.JSON(w, http.StatusBadRequest, "Format data booking tidak valid", nil)
		return
	}

	id, err := h.service.CreateBooking(req)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	booking, _ := h.service.GetBooking(id)
	response.JSON(w, http.StatusCreated, "Booking berhasil dibuat", booking)
}

// GetBooking handles GET /bookings/{id}
func (h *BookingHandler) GetBooking(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, "ID booking tidak valid", nil)
		return
	}

	booking, err := h.service.GetBooking(id)
	if err != nil {
		response.JSON(w, http.StatusNotFound, "Booking tidak ditemukan", nil)
		return
	}

	response.JSON(w, http.StatusOK, "Berhasil mendapatkan booking", booking)
}

// ListBookings handles GET /bookings
func (h *BookingHandler) ListBookings(w http.ResponseWriter, r *http.Request) {
	bookings, err := h.service.ListBookings()
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, "Gagal mendapatkan list booking", nil)
		return
	}

	response.JSON(w, http.StatusOK, "Berhasil mendapatkan list booking", bookings)
}

// CheckAvailability handles GET /bookings/check-availability?field_id=1&start_time=2025-10-25T10:00:00Z&end_time=2025-10-25T12:00:00Z
func (h *BookingHandler) CheckAvailability(w http.ResponseWriter, r *http.Request) {
	fieldIDStr := r.URL.Query().Get("field_id")
	startTimeStr := r.URL.Query().Get("start_time")
	endTimeStr := r.URL.Query().Get("end_time")

	if fieldIDStr == "" || startTimeStr == "" || endTimeStr == "" {
		response.JSON(w, http.StatusBadRequest, "Parameter field_id, start_time, dan end_time diperlukan", nil)
		return
	}

	fieldID, err := strconv.Atoi(fieldIDStr)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, "Field ID tidak valid", nil)
		return
	}

	startTime, err := parseTime(startTimeStr)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, "Format start_time tidak valid", nil)
		return
	}

	endTime, err := parseTime(endTimeStr)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, "Format end_time tidak valid", nil)
		return
	}

	available, err := h.service.CheckAvailability(fieldID, startTime, endTime)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	data := map[string]bool{
		"available": available,
	}
	response.JSON(w, http.StatusOK, "Pengecekan ketersediaan berhasil", data)
}

func parseTime(timeStr string) (time.Time, error) {
	return time.Parse(time.RFC3339, timeStr)
}

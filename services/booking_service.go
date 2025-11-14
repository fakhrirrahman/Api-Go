package services

import (
	"goApi/models"
	"goApi/repositories"
	"time"
)

type BookingService struct {
	repo repositories.BookingRepository
}

func NewBookingService(r repositories.BookingRepository) *BookingService {
	return &BookingService{repo: r}
}

func (s *BookingService) CreateBooking(req models.CreateBookingRequest) (int, error) {
	// Validate times
	if req.EndTime.Before(req.StartTime) {
		return 0, errInvalidTime
	}

	booking := models.Booking{
		FieldID:   req.FieldID,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		Status:    "pending",
	}

	return s.repo.Create(booking)
}

func (s *BookingService) GetBooking(id int) (models.Booking, error) {
	return s.repo.GetByID(id)
}

func (s *BookingService) ListBookings() ([]models.Booking, error) {
	return s.repo.GetAll()
}

func (s *BookingService) CheckAvailability(fieldID int, startTime, endTime time.Time) (bool, error) {
	overlap, err := s.repo.CheckOverlap(fieldID, startTime, endTime)
	if err != nil {
		return false, err
	}
	return !overlap, nil
}

var errInvalidTime = errCustom{message: "waktu akhir harus lebih besar dari waktu awal"}

type errCustom struct {
	message string
}

func (e errCustom) Error() string {
	return e.message
}

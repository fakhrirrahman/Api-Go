package repositories

import (
	"fmt"
	"goApi/models"
	"time"
)

type BookingRepository interface {
	Create(booking models.Booking) (int, error)
	GetByID(id int) (models.Booking, error)
	GetAll() ([]models.Booking, error)
	CheckOverlap(fieldID int, startTime, endTime time.Time) (bool, error)
	Update(booking models.Booking) error
}

type bookingRepository struct {
	data []models.Booking
}

func NewBookingRepository() BookingRepository {
	return &bookingRepository{
		data: []models.Booking{},
	}
}

func (r *bookingRepository) Create(booking models.Booking) (int, error) {
	// Check for overlapping bookings
	overlap, err := r.CheckOverlap(booking.FieldID, booking.StartTime, booking.EndTime)
	if err != nil {
		return 0, err
	}

	if overlap {
		return 0, fmt.Errorf("lapangan sudah dibooking pada waktu tersebut")
	}

	booking.ID = len(r.data) + 1
	booking.Status = "pending"
	booking.CreatedAt = time.Now()
	booking.UpdatedAt = time.Now()

	r.data = append(r.data, booking)
	return booking.ID, nil
}

func (r *bookingRepository) GetByID(id int) (models.Booking, error) {
	for _, booking := range r.data {
		if booking.ID == id {
			return booking, nil
		}
	}
	return models.Booking{}, fmt.Errorf("booking tidak ditemukan")
}

func (r *bookingRepository) GetAll() ([]models.Booking, error) {
	return r.data, nil
}

func (r *bookingRepository) CheckOverlap(fieldID int, startTime, endTime time.Time) (bool, error) {
	for _, booking := range r.data {
		// Only check active bookings (not cancelled)
		if booking.FieldID == fieldID && booking.Status != "cancelled" {
			// Check if times overlap
			if startTime.Before(booking.EndTime) && endTime.After(booking.StartTime) {
				return true, nil
			}
		}
	}
	return false, nil
}

func (r *bookingRepository) Update(booking models.Booking) error {
	for i, b := range r.data {
		if b.ID == booking.ID {
			booking.UpdatedAt = time.Now()
			r.data[i] = booking
			return nil
		}
	}
	return fmt.Errorf("booking tidak ditemukan")
}

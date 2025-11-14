package models

import "time"

type Booking struct {
	ID        int       `json:"id"`
	FieldID   int       `json:"field_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Status    string    `json:"status"` // pending, paid, cancelled
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateBookingRequest struct {
	FieldID   int       `json:"field_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

package repositories

import (
	"fmt"
	"goApi/models"
	"time"
)

type PaymentRepository interface {
	Create(payment models.Payment) (int, error)
	GetByID(id int) (models.Payment, error)
	GetByBookingID(bookingID int) (models.Payment, error)
	Update(payment models.Payment) error
}

type paymentRepository struct {
	data []models.Payment
}

func NewPaymentRepository() PaymentRepository {
	return &paymentRepository{
		data: []models.Payment{},
	}
}

func (r *paymentRepository) Create(payment models.Payment) (int, error) {
	payment.ID = len(r.data) + 1
	if payment.Status == "" {
		payment.Status = "pending"
	}
	payment.CreatedAt = time.Now().String()
	payment.UpdatedAt = time.Now().String()

	r.data = append(r.data, payment)
	return payment.ID, nil
}

func (r *paymentRepository) GetByID(id int) (models.Payment, error) {
	for _, payment := range r.data {
		if payment.ID == id {
			return payment, nil
		}
	}
	return models.Payment{}, fmt.Errorf("payment tidak ditemukan")
}

func (r *paymentRepository) GetByBookingID(bookingID int) (models.Payment, error) {
	for _, payment := range r.data {
		if payment.BookingID == bookingID {
			return payment, nil
		}
	}
	return models.Payment{}, fmt.Errorf("payment untuk booking tidak ditemukan")
}

func (r *paymentRepository) Update(payment models.Payment) error {
	for i, p := range r.data {
		if p.ID == payment.ID {
			payment.UpdatedAt = time.Now().String()
			r.data[i] = payment
			return nil
		}
	}
	return fmt.Errorf("payment tidak ditemukan")
}

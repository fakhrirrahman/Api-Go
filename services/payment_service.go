package services

import (
	"goApi/models"
	"goApi/repositories"
)

type PaymentService struct {
	paymentRepo repositories.PaymentRepository
	bookingRepo repositories.BookingRepository
}

func NewPaymentService(pr repositories.PaymentRepository, br repositories.BookingRepository) *PaymentService {
	return &PaymentService{
		paymentRepo: pr,
		bookingRepo: br,
	}
}

func (s *PaymentService) ProcessPayment(bookingID int) (models.Payment, error) {
	// Get the booking
	booking, err := s.bookingRepo.GetByID(bookingID)
	if err != nil {
		return models.Payment{}, err
	}

	// Check if payment already exists
	existingPayment, _ := s.paymentRepo.GetByBookingID(bookingID)
	if existingPayment.ID != 0 {
		return existingPayment, nil
	}

	// Create payment record
	payment := models.Payment{
		BookingID: bookingID,
		Amount:    1000, // Mock price per hour in cents
		Status:    "pending",
	}

	id, err := s.paymentRepo.Create(payment)
	if err != nil {
		return models.Payment{}, err
	}

	// Update payment status to paid (mock payment success)
	payment.ID = id
	payment.Status = "paid"
	err = s.paymentRepo.Update(payment)
	if err != nil {
		return models.Payment{}, err
	}

	// Update booking status to paid
	booking.Status = "paid"
	err = s.bookingRepo.Update(booking)
	if err != nil {
		return models.Payment{}, err
	}

	payment.ID = id
	return payment, nil
}

func (s *PaymentService) GetPayment(id int) (models.Payment, error) {
	return s.paymentRepo.GetByID(id)
}

func (s *PaymentService) GetPaymentByBooking(bookingID int) (models.Payment, error) {
	return s.paymentRepo.GetByBookingID(bookingID)
}

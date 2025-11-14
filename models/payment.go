package models

type Payment struct {
	ID        int    `json:"id"`
	BookingID int    `json:"booking_id"`
	Amount    int    `json:"amount"`
	Status    string `json:"status"` // pending, paid, failed
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ProcessPaymentRequest struct {
	BookingID int `json:"booking_id"`
}

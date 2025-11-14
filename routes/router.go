package routes

import (
	"goApi/handlers"
	"goApi/repositories"
	"goApi/services"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)
	authHandler := handlers.NewAuthHandler(userService)

	// Auth endpoints (no JWT required)
	r.HandleFunc("/auth/login", authHandler.Login).Methods("POST")
	r.HandleFunc("/auth/register", authHandler.Register).Methods("POST")
	r.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
	r.HandleFunc("/users", userHandler.CreateUser).Methods("POST")

	bookingRepo := repositories.NewBookingRepository()
	bookingService := services.NewBookingService(bookingRepo)
	bookingHandler := handlers.NewBookingHandler(bookingService)

	r.HandleFunc("/bookings", bookingHandler.ListBookings).Methods("GET")
	r.HandleFunc("/bookings", bookingHandler.CreateBooking).Methods("POST")
	r.HandleFunc("/bookings/{id}", bookingHandler.GetBooking).Methods("GET")
	r.HandleFunc("/bookings/check-availability", bookingHandler.CheckAvailability).Methods("GET")

	paymentRepo := repositories.NewPaymentRepository()
	paymentService := services.NewPaymentService(paymentRepo, bookingRepo)
	paymentHandler := handlers.NewPaymentHandler(paymentService)

	r.HandleFunc("/payments", paymentHandler.ProcessPayment).Methods("POST")
	r.HandleFunc("/payments/{id}", paymentHandler.GetPayment).Methods("GET")
	r.HandleFunc("/payments/booking/{booking_id}", paymentHandler.GetPaymentByBooking).Methods("GET")

	return r
}

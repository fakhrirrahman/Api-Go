package routes

import (
	"goApi/handlers"
	"goApi/middleware"
	"goApi/repositories"
	"goApi/services"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)
	authHandler := handlers.NewAuthHandler(userService)

	// Auth endpoints (no JWT required)
	app.Post("/auth/login", authHandler.LoginFiber)
	app.Post("/auth/register", authHandler.RegisterFiber)
	app.Get("/users/:id", userHandler.GetUserFiber)
	app.Post("/users", userHandler.CreateUserFiber)

	bookingRepo := repositories.NewBookingRepository()
	bookingService := services.NewBookingService(bookingRepo)
	bookingHandler := handlers.NewBookingHandler(bookingService)

	// Booking endpoints (with JWT middleware)
	app.Get("/bookings", middleware.JWTMiddlewareFiber, bookingHandler.ListBookingsFiber)
	app.Post("/bookings", middleware.JWTMiddlewareFiber, bookingHandler.CreateBookingFiber)
	app.Get("/bookings/:id", middleware.JWTMiddlewareFiber, bookingHandler.GetBookingFiber)
	app.Get("/bookings/check-availability", middleware.JWTMiddlewareFiber, bookingHandler.CheckAvailabilityFiber)

	paymentRepo := repositories.NewPaymentRepository()
	paymentService := services.NewPaymentService(paymentRepo, bookingRepo)
	paymentHandler := handlers.NewPaymentHandler(paymentService)

	// Payment endpoints (with JWT middleware)
	app.Post("/payments", middleware.JWTMiddlewareFiber, paymentHandler.ProcessPaymentFiber)
	app.Get("/payments/:id", middleware.JWTMiddlewareFiber, paymentHandler.GetPaymentFiber)
	app.Get("/payments/booking/:booking_id", middleware.JWTMiddlewareFiber, paymentHandler.GetPaymentByBookingFiber)
}

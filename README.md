# API Go - Booking Management System

This is a simple field booking management system built with Go, Fiber/Mux, and PostgreSQL.

## Features

- User management (Create, Read)
- Field booking with overlap checking
- Payment processing (mock implementation)
- In-memory data storage (ready for PostgreSQL migration)

## Prerequisites

- Go 1.25.3+
- PostgreSQL (optional, for production)
- Docker (optional)

## Project Structure

```
.
├── config/           # Configuration management
├── handlers/         # HTTP request handlers
├── interfaces/       # Interfaces for repositories
├── models/          # Data models
├── repositories/    # Data access layer
├── response/        # Response formatting
├── routes/          # Route definitions
├── services/        # Business logic
├── migrations/      # Database migrations
└── main.go         # Entry point
```

## Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd Api-Go
```

2. Install dependencies:
```bash
go mod download
```

3. Create `.env` file:
```env
APP_NAME=BookingAPI
APP_PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=booking_db
JWT_SECRET=your_secret_key
```

4. Run the server:
```bash
go run main.go
```

The server will start at `http://localhost:8080`

## API Endpoints

### Users
- `GET /users` - List all users
- `GET /users/{id}` - Get user by ID
- `POST /users` - Create a new user

### Bookings
- `GET /bookings` - List all bookings
- `GET /bookings/{id}` - Get booking by ID
- `POST /bookings` - Create a new booking
- `GET /bookings/check-availability?field_id=1&start_time=2025-10-25T10:00:00Z&end_time=2025-10-25T12:00:00Z` - Check field availability

### Payments
- `POST /payments` - Process payment for a booking
- `GET /payments/{id}` - Get payment by ID
- `GET /payments/booking/{booking_id}` - Get payment by booking ID

## API Usage Examples

### Create a Booking
```bash
curl -X POST http://localhost:8080/bookings \
  -H "Content-Type: application/json" \
  -d '{
    "field_id": 1,
    "start_time": "2025-10-25T10:00:00Z",
    "end_time": "2025-10-25T12:00:00Z"
  }'
```

### Check Availability
```bash
curl "http://localhost:8080/bookings/check-availability?field_id=1&start_time=2025-10-25T10:00:00Z&end_time=2025-10-25T12:00:00Z"
```

### Process Payment
```bash
curl -X POST http://localhost:8080/payments \
  -H "Content-Type: application/json" \
  -d '{
    "booking_id": 1
  }'
```

## Docker Setup

Build the Docker image:
```bash
docker build -t booking-api .
```

Run with Docker Compose:
```bash
docker-compose up
```

## Database Migration

To set up PostgreSQL:

1. Create database:
```bash
createdb booking_db
```

2. Run migrations:
```bash
psql booking_db < migrations/001_initial_schema.sql
```

## Development

To run tests:
```bash
go test ./...
```

## Future Enhancements

- [ ] Implement JWT authentication
- [ ] Add PostgreSQL integration
- [ ] Add Swagger documentation
- [ ] Add unit tests
- [ ] Add field management endpoints
- [ ] Add cancellation feature
- [ ] Add email notifications

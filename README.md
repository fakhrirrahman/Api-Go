````markdown
# 📱 Booking API - Golang + Fiber + PostgreSQL

Sistem API untuk booking lapangan dengan authentication JWT dan payment mock.

## ✨ Features

- ✅ **JWT Authentication** - Login & Register dengan token 24 jam
- ✅ **Booking System** - Create, Read, List bookings dengan overlap check
- ✅ **Payment Processing** - Mock payment API yang update booking status
- ✅ **PostgreSQL Integration** - Database persistence
- ✅ **Fiber Framework** - Fast & modern Go web framework
- ✅ **Docker Support** - Easy deployment

## 🛠️ Tech Stack

- **Language**: Go 1.25.3
- **Framework**: Fiber v2
- **Database**: PostgreSQL
- **Authentication**: JWT
- **Containerization**: Docker

## 🚀 Quick Start

### Option 1: Dengan Docker (Recommended)

```bash
# 1. Navigate to project
cd /home/killua/project/Api-Go

# 2. Run setup script
bash quick_setup.sh

# 3. Start server
./goApi
```

### Option 2: Manual PostgreSQL Setup

```bash
# 1. Setup database
bash setup_manual.sh

# 2. Build & run
go build -o goApi
./goApi
```

## 📚 API Endpoints

### Authentication
- `POST /auth/login` - Login & dapatkan JWT token
- `POST /auth/register` - Register user baru

### Bookings (Perlu JWT)
- `POST /bookings` - Create booking
- `GET /bookings` - List semua booking
- `GET /bookings/:id` - Get detail booking
- `GET /bookings/check-availability` - Cek ketersediaan lapangan

### Payments (Perlu JWT)
- `POST /payments` - Proses payment
- `GET /payments/:id` - Get payment detail
- `GET /payments/booking/:booking_id` - Get payment by booking

## 🧪 Testing

### 1. Login

```bash
curl -X POST http://localhost:8080/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"user@test.com","password":"pass"}'
```

### 2. Create Booking

```bash
curl -X POST http://localhost:8080/bookings \
  -H "Authorization: Bearer <TOKEN>" \
  -H "Content-Type: application/json" \
  -d '{
    "field_id": 1,
    "start_time": "2025-10-25T10:00:00Z",
    "end_time": "2025-10-25T12:00:00Z"
  }'
```

### 3. Process Payment

```bash
curl -X POST http://localhost:8080/payments \
  -H "Authorization: Bearer <TOKEN>" \
  -H "Content-Type: application/json" \
  -d '{"booking_id": 1}'
```

## 📁 Project Structure

```
Api-Go/
├── main.go                    # Entry point
├── config/                    # Configuration
├── database/                  # Database connection
├── models/                    # Data models
├── repositories/              # Data access layer
├── services/                  # Business logic
├── handlers/                  # HTTP handlers
├── routes/                    # Route definitions
├── middleware/                # Middleware (JWT)
├── utils/                     # Utilities
├── migrations/                # Database migrations
├── docker-compose.yml         # Docker setup
├── Dockerfile                 # Container config
└── README.md                  # This file
```

## 🗄️ Database

Tables:
- `users` - User data
- `fields` - Lapangan data
- `bookings` - Booking data (dengan overlap check)
- `payments` - Payment data

## 📝 Environment Variables (.env)

```
APP_NAME=BookingAPI
APP_PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=1
DB_NAME=booking_db
JWT_SECRET=your_secret_key_here
```

## 📚 Additional Documentation

- `DATABASE_QUICK_START.md` - Database setup cepat
- `DATABASE_SETUP.md` - Database setup lengkap
- `JWT_QUICK_START.md` - JWT quick guide
- `JWT_GUIDE.md` - JWT dokumentasi lengkap
- `API_Postman_Collection.json` - Postman collection

## 📦 Dependencies

```
github.com/gofiber/fiber/v2    - Web framework
github.com/golang-jwt/jwt/v5   - JWT authentication
github.com/lib/pq               - PostgreSQL driver
github.com/joho/godotenv        - Environment loader
```

## ✅ Requirement Checklist

- ✅ Golang + Fiber framework
- ✅ PostgreSQL database integration
- ✅ JWT authentication
- ✅ Booking endpoint with overlap check
- ✅ Payment mock API
- ✅ Docker support
- ✅ Postman collection

---

**Status: ✨ Production Ready**
````

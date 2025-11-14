# 🔐 JWT Authentication Guide

## Cara Mendapatkan JWT Token

### 1. **Login Endpoint** (Dapatkan Token)

**URL:** `POST /auth/login`

**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "password123"
}
```

**Response (Success):**
```json
{
  "status": 200,
  "message": "Login berhasil",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "ID": 1,
      "Name": "User",
      "Email": "user@example.com"
    },
    "expiry": "24 jam"
  }
}
```

### 2. **Register Endpoint** (Buat Akun & Dapatkan Token)

**URL:** `POST /auth/register`

**Request Body:**
```json
{
  "ID": 0,
  "Name": "John Doe",
  "Email": "john@example.com"
}
```

**Response:**
```json
{
  "status": 201,
  "message": "Register berhasil",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "ID": 1,
      "Name": "John Doe",
      "Email": "john@example.com"
    },
    "expiry": "24 jam"
  }
}
```

---

## 🎯 Menggunakan JWT Token

Setelah mendapatkan token dari login/register, gunakan token tersebut di header `Authorization` untuk setiap request yang memerlukan autentikasi:

```
Authorization: Bearer <JWT_TOKEN>
```

### Contoh Request dengan Token:

```bash
curl -X POST http://localhost:8080/bookings \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "Content-Type: application/json" \
  -d '{
    "field_id": 1,
    "start_time": "2025-10-25T10:00:00Z",
    "end_time": "2025-10-25T12:00:00Z"
  }'
```

---

## 📋 Token Payload

Token JWT mengandung informasi berikut:

```json
{
  "user_id": 1,
  "email": "user@example.com",
  "name": "User",
  "exp": 1634567890,    // Waktu kadaluarsa (24 jam dari dibuat)
  "iat": 1634481490,    // Waktu dibuat
  "nbf": 1634481490     // Waktu mulai berlaku
}
```

---

## ⏰ Token Expiry

- **Default Duration:** 24 jam
- **Setelah kadaluarsa:** User harus login ulang

---

## 🔒 Environment Variable

Set JWT_SECRET di `.env` file untuk keamanan lebih:

```
JWT_SECRET=your-super-secret-key-here
```

Jika tidak di-set, akan menggunakan default secret.

---

## ✅ Endpoints yang Memerlukan JWT

- `POST /bookings` - Create booking
- `GET /bookings` - List bookings
- `GET /bookings/{id}` - Get booking detail
- `GET /bookings/check-availability` - Check availability
- `POST /payments` - Process payment
- `GET /payments/{id}` - Get payment detail
- `GET /payments/booking/{booking_id}` - Get payment by booking

## ❌ Endpoints Tanpa JWT

- `POST /auth/login` - Login
- `POST /auth/register` - Register
- `GET /users/{id}` - Get user
- `POST /users` - Create user

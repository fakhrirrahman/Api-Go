# 📌 JWT Token - Ringkas

## **Cara Mendapatkan JWT Token**

JWT token didapatkan dari endpoint **login/register**:

### **Option 1: Login**
```bash
POST /auth/login
Body: {
  "email": "user@example.com",
  "password": "password123"
}
```

**Response:**
```json
{
  "status": 200,
  "message": "Login berhasil",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {...},
    "expiry": "24 jam"
  }
}
```

### **Option 2: Register**
```bash
POST /auth/register
Body: {
  "Name": "John Doe",
  "Email": "john@example.com"
}
```

---

## **Menggunakan Token**

Semua request ke endpoint yang membutuhkan auth harus include token di header:

```bash
Authorization: Bearer <TOKEN_DARI_LOGIN>
```

### **Contoh:**
```bash
curl -X POST http://localhost:8080/bookings \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "Content-Type: application/json" \
  -d '{"field_id": 1, "start_time": "...", "end_time": "..."}'
```

---

## **Token Info**

- ✅ **Berlaku:** 24 jam
- ✅ **Cara dapat:** Login/Register
- ✅ **Tempat pakai:** Header Authorization

---

## **Quick Test di Postman**

1. Import file: `API_Postman_Collection.json`
2. Jalankan request "Login" atau "Register"
3. Token otomatis tersimpan di variable `{{jwt_token}}`
4. Gunakan di request lainnya

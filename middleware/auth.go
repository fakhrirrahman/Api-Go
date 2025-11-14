package middleware

import (
	"goApi/response"
	"goApi/utils"
	"net/http"
	"strings"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			response.JSON(w, http.StatusUnauthorized, "Authorization header diperlukan", nil)
			return
		}

		// Extract token from "Bearer <token>"
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			response.JSON(w, http.StatusUnauthorized, "Format authorization tidak valid", nil)
			return
		}

		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			response.JSON(w, http.StatusUnauthorized, "Token tidak valid: "+err.Error(), nil)
			return
		}

		// Store claims in context for use in handlers
		r.Header.Set("X-User-ID", string(rune(claims.UserID)))
		r.Header.Set("X-User-Email", claims.Email)
		r.Header.Set("X-User-Name", claims.Name)

		next.ServeHTTP(w, r)
	})
}

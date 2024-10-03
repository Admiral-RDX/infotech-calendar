package auth_helpers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

var JWT_SECRET = []byte(os.Getenv("JWT_SECRET"))

func HasValidSession(r *http.Request) bool {
	cookie, err := r.Cookie("session_token")
	if err == nil && cookie.Value != "" {
		return true
	}

	return false
}

func GenerateJWTToken(userID string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JWT_SECRET)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWTToken(tokenString string) bool {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return JWT_SECRET, nil
	})

	if err != nil {
		return false
	}

	return token.Valid
}

package infrastructure

import (
	"log"
	"os"
	"task_manager/domain"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

var GenerateJWT = func(user *domain.User) (string, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
    jwtKey := os.Getenv("JWTKEY")
    if jwtKey == "" {
        log.Fatal("JWTKEY is not set in .env file")
    }

	// Set expiration time for the token
	expirationTime := time.Now().Add(24 * time.Hour).Unix()

	// Define custom claims structure
	claims := &domain.Claims{
		ID:       user.ID,
		Username: user.Username,
		Role:     user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
		},
	}

	// Create a new JWT token with the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the JWT key
	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

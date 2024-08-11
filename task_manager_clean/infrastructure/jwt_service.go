package infrastructure

import (
	"errors"
	"os"
	"time"
	"task_manager/domain"

	"github.com/golang-jwt/jwt"
)

func GenerateJWT(user *domain.User) (string, error) {
	jwtKey, ok := os.LookupEnv("JWT_KEY")
	if !ok {
		return "", errors.New("JWT_KEY not found")
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

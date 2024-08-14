package infrastructure

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strings"
	"task_manager/domain"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func ValidateJWT(tokenString string) (*domain.Claims, error) {
    if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
    jwtKey := os.Getenv("JWTKEY")
    if jwtKey == "" {
        log.Fatal("JWT_KEY not found in environment")
    }

    jwtKeyBytes := []byte(jwtKey)

    // Parse the token with the claims
    claims := &domain.Claims{}
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtKeyBytes, nil
    })
    if err != nil {
        if err == jwt.ErrSignatureInvalid {
            return nil, errors.New("invalid signature")
        }
        return nil, err
    }
    if !token.Valid {
        return nil, errors.New("validatejwt func invalid token")
    }
    return claims, nil
}

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
            c.Abort()
            return
        }

        tokenString := strings.TrimPrefix(authHeader,"Bearer ")

        user, err := ValidateJWT(tokenString)
        if err != nil {
            c.JSON(http.StatusUnauthorized, err.Error())
            c.Abort()
            return
        }

        c.Set("user", user)
        c.Next()
    }
}
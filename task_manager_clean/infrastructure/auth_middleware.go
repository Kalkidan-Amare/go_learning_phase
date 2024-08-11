package infrastructure

import (
	"errors"
	"net/http"
	"os"
	"strings"
	"task_manager/domain"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func ValidateJWT(tokenString string) (*domain.Claims, error) {
    jwtKey, ok := os.LookupEnv("JWT_KEY")
	if !ok {
		return nil, errors.New("JWT_KEY not found in environment")
	}

    claims := &domain.Claims{}
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
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
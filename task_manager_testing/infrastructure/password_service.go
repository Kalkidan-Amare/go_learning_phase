package infrastructure

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

var HashPassword = func(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		log.Fatal("Internal server error")
	}

	return string(hashedPassword), nil
}

var ComparePassword = func(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("could not hash password")
	}
	return string(hashedPassword), nil
}

func VerifyPassword(HashedPassword string, candidatePassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(HashedPassword), []byte(candidatePassword))
}

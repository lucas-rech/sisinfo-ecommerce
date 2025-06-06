package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password *string) (string, error) {
	if password == nil || *password == "" {
		return "", fmt.Errorf("password cannot be empty")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error hashing password: %w", err)
	}

	return string(hashedPassword), nil
}

func CheckPasswordHash(password, hash *string) error {
	if password == nil || hash == nil || *password == "" || *hash == "" {
		return fmt.Errorf("password and hash cannot be empty")
	}

	err := bcrypt.CompareHashAndPassword([]byte(*hash), []byte(*password))
	if err != nil {
		return fmt.Errorf("password does not match hash: %w", err)
	}

	return nil 
}

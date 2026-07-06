package utils

import (
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword meng-hash password menggunakan bcrypt
func HashPassword(password string) (string, error) {

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

// ComparePassword membandingkan password plain dengan hash yang tersimpan
func ComparePassword(hashed, plain string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))

	return err == nil
}

// GenerateRandomToken menghasilkan token random, berguna untuk password reset
func GenerateRandomToken(length int) (string, error) {

	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(bytes), nil
}

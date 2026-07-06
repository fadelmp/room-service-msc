package utils

import (
	"net/mail"
	"regexp"
)

var phoneRegex = regexp.MustCompile(`^(\+62|62|0)8[1-9][0-9]{6,10}$`)

// IsValidEmail mengecek format email secara sintaks (bukan validasi keberadaan/uniqueness)
func IsValidEmail(email string) bool {

	_, err := mail.ParseAddress(email)

	return err == nil
}

// IsValidPhoneNumber mengecek format nomor HP Indonesia
// Menerima format: 08xxx, 62xxx, +62xxx
func IsValidPhoneNumber(phone string) bool {

	return phoneRegex.MatchString(phone)
}

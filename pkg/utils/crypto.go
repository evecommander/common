package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// BcryptEncode encode string to bcrypt crypto string
func BcryptEncode(originstr string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(originstr), bcrypt.DefaultCost)

	if err != nil {
		return ""
	}

	return string(bytes)
}

// BcryptCheck compare password and hash
func BcryptCheck(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

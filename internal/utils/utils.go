package utils

import (
	"os"
	"github.com/codehack/scrypto"
)

// GetEnvOrDefault return a env value or a default if not exists
func GetEnvOrDefault(key, d string) string {
	v := os.Getenv(key)
	if v == "" {
		return d
	}
	return v
}

func HashPassword(password string)(string, error) {
	return scrypto.Hash(password)
}

func CheckPassword(pass, hashed string) bool {
	return scrypto.Compare(pass, hashed)
}

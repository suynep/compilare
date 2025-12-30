package crypt

import (
	"crypto/sha256"
	"fmt"
)

func GenerateSessionKey(salt string) {

}

func HashPassword(password string) string {
	// placeholder for the time being; actual logic shall be implemented soon
	h := sha256.New()
	h.Write([]byte(password))

	bs := h.Sum(nil)

	return fmt.Sprintf("%x", bs)
}

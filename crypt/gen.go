package crypt

import (
	"crypto/sha256"
	"fmt"

	"github.com/suynep/compilare/database"
	"github.com/suynep/compilare/types"
)

func GenerateSessionKey(salt string) {

}

func HashPassword(password string) string {
	// placeholder for the time being; actual logic shall be implemented soon
	// Jan 1, '26 : this func needs a secure treatment (a better algorithm)
	h := sha256.New()
	h.Write([]byte(password))

	bs := h.Sum(nil)

	return fmt.Sprintf("%x", bs)
}

func CheckPassword(user types.LoginUser) (bool, error) {
	regUser, err := database.GetUserByUsername(user.Username)

	if err != nil {
		return false, err
	}

	if HashPassword(user.Password) == regUser.Password {
		return true, nil
	}

	return false, nil
}

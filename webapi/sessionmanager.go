package webapi

import (
	"log"
	"time"

	"github.com/suynep/compilare/database"
	"github.com/suynep/compilare/types"
	"github.com/suynep/compilare/webapi/utils"
)

func CheckSessionExpiration(session types.Session) (bool, error) {
	/*
		This should be run in a short-poll fashion, perhaps every 2 seconds? (USE GOROUTINES!!!)
	*/
	startTime := session.CreatedAt

	if time.Since(startTime).Minutes() >= utils.SESSION_EXPIRATION_DELTA {
		return true, nil
	}

	return false, nil
}

func SessionPopper(session types.Session) {
	for {
		hasExpired, _ := CheckSessionExpiration(session)
		if hasExpired {
			_ = database.RemoveSession(session)
			log.Printf("Time exceeded %.2f minutes! Removing Session...\n", utils.SESSION_EXPIRATION_DELTA)
			log.Printf("Session %s removed successfully!\n", session.SessionKey)
			break
		}
	}
}

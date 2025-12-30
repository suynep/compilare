package webapi

import (
	"time"

	"github.com/suynep/compilare/types"
	"github.com/suynep/compilare/webapi/utils"
)

func CheckSessionExpiration(session types.Session) (bool, error) {
	/*
		This should be run in a short-poll fashion, perhaps every 2 seconds?
	*/
	startTime := session.CreatedAt

	if time.Since(startTime).Minutes() >= utils.SESSION_EXPIRATION_DELTA {
		return true, nil
	}

	return false, nil
}

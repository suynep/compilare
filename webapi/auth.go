package webapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/suynep/compilare/database"
	"github.com/suynep/compilare/types"
)

const COOKIE_NAME = "GO_SESSION_ID"

func RegistrationHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		credentials, err := io.ReadAll(r.Body)

		if err != nil {
			response := map[string]string{"message": "No credentials provided"}
			json.NewEncoder(w).Encode(response)
			return
		}

		regUser := new(types.RegisterUser)
		err = json.Unmarshal(credentials, regUser)

		if err != nil {
			response := map[string]string{"message": "Unparseable credentials provided"}
			json.NewEncoder(w).Encode(response)
			return
		}

		u_id := database.InsertUser(*regUser)

		newSession := types.Session{
			SessionKey: uuid.NewString(),
			UserId:     int(u_id),
		}

		_ = database.InsertSession(newSession)

		w.Header().Add("Set-Cookie", fmt.Sprintf("%s=%s", COOKIE_NAME, newSession.SessionKey))

		response := map[string]string{
			"message": "Registration Successful!",
		}

		json.NewEncoder(w).Encode(response)

	} else if r.Method == "GET" {
		response := map[string]string{
			"message": "Welcome to Register Route! Make a POST request to register a new user.",
		}

		json.NewEncoder(w).Encode(response)
	} else {
		response := map[string]string{
			"message": "Method not Supported. Yet.",
		}

		json.NewEncoder(w).Encode(response)

	}
}

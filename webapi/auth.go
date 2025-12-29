package webapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/suynep/compilare/types"
)

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

	}

}

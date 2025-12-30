package webapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/suynep/compilare/crypt"
	"github.com/suynep/compilare/database"
	"github.com/suynep/compilare/types"
)

const COOKIE_NAME = "GO_SESSION_ID"

func RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	if r.Method == "POST" {
		credentials, err := io.ReadAll(r.Body)

		if err != nil {
			response := map[string]string{"message": "No credentials provided"}
			json.NewEncoder(w).Encode(response)
			return
		}

		regUser := new(types.RegisterUser)
		err = json.Unmarshal(credentials, regUser)
		regUser.Password = crypt.HashPassword(regUser.Password) // hash pw (unimplemented as of now)

		if err != nil {
			response := map[string]string{"message": "Unparseable credentials provided"}
			json.NewEncoder(w).Encode(response)
			return
		}

		_, err = database.InsertUser(*regUser)

		if err != nil {
			response := map[string]string{"message": "Username or Email already exists"}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return

		}

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

		w.WriteHeader(http.StatusMethodNotAllowed)

		json.NewEncoder(w).Encode(response)

	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	if r.Method == "POST" {
		credentials, err := io.ReadAll(r.Body)

		if err != nil {
			response := map[string]string{"message": "No credentials provided"}
			json.NewEncoder(w).Encode(response)
			return
		}
		loginUser := new(types.LoginUser)

		err = json.Unmarshal(credentials, loginUser)

		if err != nil {
			log.Fatalf("Error while parsing data: %v", err)
		}

		regUser, err := database.GetUserByUsername(loginUser.Username)

		if err != nil {
			response := map[string]string{"message": "Login Failed"}
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(response)
			return
		}

		check, err := crypt.CheckPassword(*loginUser)

		if err != nil {
			response := map[string]string{"message": "Login Failed"}
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(response)
			return
		}

		if check {
			newSession := types.Session{
				SessionKey: uuid.NewString(),
				UserId:     int(regUser.Id),
			}

			_ = database.InsertSession(newSession)
			w.Header().Add("Set-Cookie", fmt.Sprintf("%s=%s", COOKIE_NAME, newSession.SessionKey))
			response := map[string]string{"message": "Login Successful", "session_key": newSession.SessionKey}

			json.NewEncoder(w).Encode(response)

		} else {
			response := map[string]string{"message": "Login Failed"}
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(response)
			return
		}
	} else if r.Method == "GET" {
		response := map[string]string{
			"message": "Welcome to login route. Make a POST request to this endpoint (username, password) to log in to the application!",
		}
		json.NewEncoder(w).Encode(response)

	} else {
		response := map[string]string{
			"message": "Method not Supported. Yet.",
		}

		w.WriteHeader(http.StatusMethodNotAllowed)

		json.NewEncoder(w).Encode(response)
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	if r.Method == "GET" {
		if r.Header.Get("X-User") != "" {
			user, err := database.GetUserByUsername(r.Header.Get("X-User"))
			if err != nil {
				response := map[string]string{
					"message": "Error while fetching User from DB",
				}
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(response)
			}

			sessionKey, err := r.Cookie("GO_SESSION_ID")
			if err != nil {
				response := map[string]string{
					"message": "No authentication details provided!",
				}
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(response)
			}
			session := types.Session{
				SessionKey: sessionKey.Value,
				UserId:     int(user.Id),
			}

			err = database.RemoveSession(session)

			if err != nil {
				log.Printf("Error while deleting session from DB")
			}

			response := map[string]string{
				"message": "Logout Successful",
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(response)

		} else {
			response := map[string]string{
				"message": "Error while logging out",
			}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
		}
	}
}

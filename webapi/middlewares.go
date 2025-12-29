package webapi

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/suynep/compilare/database"
)

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("[%s] %s %s %s\n", r.Method, r.URL.Path, time.Since(startTime), r.RemoteAddr)
	})
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		goSessionId, err := r.Cookie("GO_SESSION_ID")
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Error, no authentication details provided",
			})
			return
		}

		log.Printf("%s\n", goSessionId.Value)
		session, err := database.GetSessionByKey(goSessionId.Value)
		log.Printf("%s\n", session.SessionKey)

		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Invalid Authentication Credentials",
			})
			return
		}

		user, err := database.GetUserById(int64(session.UserId))

		r.Header.Add("X-User", user.Username)
		w.Header().Add("X-User", user.Username)

		next.ServeHTTP(w, r)

		log.Printf("Session Id Received: %s", goSessionId.Value)
	})
}

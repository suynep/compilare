package webapi

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
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
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Error, no authentication details provided",
			})
			return
		}

		next.ServeHTTP(w, r)

		log.Printf("Session Id Received: %s", goSessionId.Value)
	})
}

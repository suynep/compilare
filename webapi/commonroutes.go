package webapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// root(/) handler
func RootHandler(w http.ResponseWriter, r *http.Request) {
	trimmed := strings.TrimRight(r.URL.Path, "/")
	if trimmed == "" {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)

		response := map[string]any{
			"message": fmt.Sprintf("Welcome to %s route! If you're trying to explore this API's functionalities, take a look at /info/universe", r.URL.Path),
			"path":    r.URL.Path,
		}
		json.NewEncoder(w).Encode(response)
	}
}

func InfoUniverseHandler(w http.ResponseWriter, r *http.Request) {
	trimmed := strings.TrimRight(r.URL.Path, "/")
	if trimmed == "/info/universe" {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)

		response := map[string]any{
			"message": fmt.Sprintf("Welcome to %s route! The functionalities of this api are list under \"functionalities\"", r.URL.Path),
			"path":    r.URL.Path,
			"functionalities": map[string]string{
				"/fetch/new/":    "fetch new posts from hackernews",
				"/fetch/top/":    "fetch top posts from hackernews",
				"/fetch/best/":   "fetch best posts from hackernews",
				"/fetch/aeon/":   "fetch posts from aeon.co",
				"/fetch/psyche/": "fetch posts from psyche.co",
			},
		}
		json.NewEncoder(w).Encode(response)
	}
}

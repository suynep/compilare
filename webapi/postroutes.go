package webapi

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/suynep/compilare/database"
)

func FetchTopHackernewsPosts(w http.ResponseWriter, r *http.Request) {
	trimmed := strings.TrimRight(r.URL.Path, "/")
	if trimmed == "/fetch/top" {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		posts := database.ReadForMemoization("t")

		response := map[string]any{
			"message": "ok",
			"path":    trimmed,
			"body":    posts,
		}
		json.NewEncoder(w).Encode(response)
	}
}

func FetchNewHackernewsPosts(w http.ResponseWriter, r *http.Request) {
	trimmed := strings.TrimRight(r.URL.Path, "/")
	if trimmed == "/fetch/new" {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		posts := database.ReadForMemoization("n")

		response := map[string]any{
			"message": "ok",
			"path":    trimmed,
			"body":    posts,
		}
		json.NewEncoder(w).Encode(response)
	}
}

func FetchBestHackernewsPosts(w http.ResponseWriter, r *http.Request) {
	trimmed := strings.TrimRight(r.URL.Path, "/")
	if trimmed == "/fetch/best" {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		posts := database.ReadForMemoization("b")

		response := map[string]any{
			"message": "ok",
			"path":    trimmed,
			"body":    posts,
		}
		json.NewEncoder(w).Encode(response)
	}
}

func FetchAeonPosts(w http.ResponseWriter, r *http.Request) {
	trimmed := strings.TrimRight(r.URL.Path, "/")

	if trimmed == "/fetch/aeon" {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		posts := database.ReadAeonPosts()

		response := map[string]any{
			"message": "ok",
			"path":    trimmed,
			"body":    posts,
		}
		json.NewEncoder(w).Encode(response)
	}
}

func FetchPsychePosts(w http.ResponseWriter, r *http.Request) {
	trimmed := strings.TrimRight(r.URL.Path, "/")

	if trimmed == "/fetch/psyche" {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		posts := database.ReadPsychePosts()

		response := map[string]any{
			"message": "ok",
			"path":    trimmed,
			"body":    posts,
		}
		json.NewEncoder(w).Encode(response)
	}
}

func AuthTestRoute(w http.ResponseWriter, r *http.Request) {
	trimmed := strings.TrimRight(r.URL.Path, "/")
	if trimmed == "/test/auth" {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)

		response := map[string]any{
			"message": "Authentication Successful!",
		}

		json.NewEncoder(w).Encode(response)
	}
}

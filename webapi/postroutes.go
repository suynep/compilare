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
		posts := database.ReadPsychePosts()

		response := map[string]any{
			"message": "ok",
			"path":    trimmed,
			"body":    posts,
		}
		json.NewEncoder(w).Encode(response)
	}
}

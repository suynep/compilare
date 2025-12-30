package webapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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

func FavoriteHackernewsPost(w http.ResponseWriter, r *http.Request) {
	trimmed := strings.TrimRight(r.URL.Path, "/")

	if strings.HasPrefix(trimmed, "/favorite/hackernews/") {
		strId := strings.TrimPrefix(trimmed, "/favorite/hackernews/")
		id, err := strconv.Atoi(strId)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := map[string]string{
				"message": fmt.Sprintf("%s is not a valid integer", strId),
			}
			json.NewEncoder(w).Encode(response)
			return
		}

		// username := r.Header.Get("X-User")
		sessionKey, err := r.Cookie("GO_SESSION_ID")

		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			response := map[string]string{
				"message": "No Authentication",
			}
			json.NewEncoder(w).Encode(response)
			return
		}

		session, err := database.GetSessionByKey(sessionKey.Value)

		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			response := map[string]string{
				"message": "No such session!",
			}
			json.NewEncoder(w).Encode(response)
			return
		}

		currentUser, err := database.GetUserById(int64(session.UserId))

		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			response := map[string]string{
				"message": "Invalid Credentials",
			}
			json.NewEncoder(w).Encode(response)
			return
		}

		if err = database.FavoriteHackernewsPost(int64(id), currentUser); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := map[string]string{
				"message": "Error while inserting favorites to DB",
			}
			json.NewEncoder(w).Encode(response)
			return
		}

		w.WriteHeader(http.StatusOK)
		response := map[string]string{
			"message": "Favorite successful",
		}
		json.NewEncoder(w).Encode(response)
	}
}

func FavoritePsychePost(w http.ResponseWriter, r *http.Request) {
	trimmed := strings.TrimRight(r.URL.Path, "/")

	if strings.HasPrefix(trimmed, "/favorite/psyche/") {
		strId := strings.TrimPrefix(trimmed, "/favorite/psyche/")
		id, err := strconv.Atoi(strId)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := map[string]string{
				"message": fmt.Sprintf("%s is not a valid integer", strId),
			}
			json.NewEncoder(w).Encode(response)
			return
		}

		// username := r.Header.Get("X-User")
		sessionKey, err := r.Cookie("GO_SESSION_ID")

		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			response := map[string]string{
				"message": "No Authentication",
			}
			json.NewEncoder(w).Encode(response)
			return
		}

		session, err := database.GetSessionByKey(sessionKey.Value)

		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			response := map[string]string{
				"message": "No such session!",
			}
			json.NewEncoder(w).Encode(response)
			return
		}

		currentUser, err := database.GetUserById(int64(session.UserId))

		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			response := map[string]string{
				"message": "Invalid Credentials",
			}
			json.NewEncoder(w).Encode(response)
			return
		}

		if err = database.FavoritePsychePost(int64(id), currentUser); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := map[string]string{
				"message": "Error while inserting favorites to DB",
			}
			json.NewEncoder(w).Encode(response)
			return
		}
		w.WriteHeader(http.StatusOK)
		response := map[string]string{
			"message": "Favorite successful",
		}
		json.NewEncoder(w).Encode(response)
	}
}

func FavoriteAeonPost(w http.ResponseWriter, r *http.Request) {
	trimmed := strings.TrimRight(r.URL.Path, "/")
	if strings.HasPrefix(trimmed, "/favorite/aeon/") {
		strId := strings.TrimPrefix(trimmed, "/favorite/aeon/")
		id, err := strconv.Atoi(strId)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := map[string]string{
				"message": fmt.Sprintf("%s is not a valid integer", strId),
			}
			json.NewEncoder(w).Encode(response)
			return
		}

		// username := r.Header.Get("X-User")
		sessionKey, err := r.Cookie("GO_SESSION_ID")

		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			response := map[string]string{
				"message": "No Authentication",
			}
			json.NewEncoder(w).Encode(response)
			return
		}

		session, err := database.GetSessionByKey(sessionKey.Value)

		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			response := map[string]string{
				"message": "No such session!",
			}
			json.NewEncoder(w).Encode(response)
			return
		}

		currentUser, err := database.GetUserById(int64(session.UserId))

		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			response := map[string]string{
				"message": "Invalid Credentials",
			}
			json.NewEncoder(w).Encode(response)
			return
		}

		if err = database.FavoriteAeonPost(int64(id), currentUser); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := map[string]string{
				"message": "Error while inserting favorites to DB",
			}
			json.NewEncoder(w).Encode(response)
			return
		}
		w.WriteHeader(http.StatusOK)
		response := map[string]string{
			"message": "Favorite successful",
		}
		json.NewEncoder(w).Encode(response)
	}
}

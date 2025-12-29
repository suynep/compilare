package webapi

import (
	"log"
	"net/http"
)

func ConnectHandlers() {
	http.Handle("/", LogMiddleware(http.HandlerFunc(RootHandler)))
	http.Handle("/info/universe", LogMiddleware(http.HandlerFunc(InfoUniverseHandler)))
	http.Handle("/info/universe/", LogMiddleware(http.HandlerFunc(InfoUniverseHandler)))
	http.Handle("/fetch/top", LogMiddleware(http.HandlerFunc(FetchTopHackernewsPosts)))
	http.Handle("/fetch/top/", LogMiddleware(http.HandlerFunc(FetchTopHackernewsPosts)))
	http.Handle("/fetch/new/", LogMiddleware(http.HandlerFunc(FetchNewHackernewsPosts)))
	http.Handle("/fetch/new", LogMiddleware(http.HandlerFunc(FetchNewHackernewsPosts)))
	http.Handle("/fetch/best/", LogMiddleware(http.HandlerFunc(FetchBestHackernewsPosts)))
	http.Handle("/fetch/best", LogMiddleware(http.HandlerFunc(FetchBestHackernewsPosts)))
	http.Handle("/fetch/aeon", LogMiddleware(http.HandlerFunc(FetchAeonPosts)))
	http.Handle("/fetch/aeon/", LogMiddleware(http.HandlerFunc(FetchAeonPosts)))
	http.Handle("/fetch/psyche", LogMiddleware(http.HandlerFunc(FetchPsychePosts)))
	http.Handle("/fetch/psyche/", LogMiddleware(http.HandlerFunc(FetchPsychePosts)))
}

func Serve() {
	ConnectHandlers()
	if err := http.ListenAndServe(":61666", nil); err != nil {
		log.Fatalf("Error occurred while starting http server: %v", err)
	}
}

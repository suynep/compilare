package webapi

import (
	"fmt"
	"log"
	"net/http"
)

const PORT = `61666`

func ConnectHandlers() {
	http.Handle("/", LogMiddleware(http.HandlerFunc(RootHandler)))
	// there MUST be a succint way to do the following w/o redundancy
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

	http.Handle("/auth/register", LogMiddleware(http.HandlerFunc(RegistrationHandler)))
	http.Handle("/auth/register/", LogMiddleware(http.HandlerFunc(RegistrationHandler)))

	http.Handle("/auth/login", LogMiddleware(http.HandlerFunc(LoginHandler)))
	http.Handle("/auth/login/", LogMiddleware(http.HandlerFunc(LoginHandler)))

	http.Handle("/auth/logout", LogMiddleware(AuthMiddleware(http.HandlerFunc(LogoutHandler))))
	http.Handle("/auth/logout/", LogMiddleware(AuthMiddleware(http.HandlerFunc(LogoutHandler))))

	http.Handle("/favorite/hackernews/", LogMiddleware(AuthMiddleware(http.HandlerFunc(FavoriteHackernewsPost))))
	http.Handle("/favorite/hackernews", LogMiddleware(AuthMiddleware(http.HandlerFunc(FavoriteHackernewsPost))))

	http.Handle("/favorite/aeon/", LogMiddleware(AuthMiddleware(http.HandlerFunc(FavoriteAeonPost))))
	http.Handle("/favorite/aeon", LogMiddleware(AuthMiddleware(http.HandlerFunc(FavoriteAeonPost))))

	http.Handle("/favorite/psyche/", LogMiddleware(AuthMiddleware(http.HandlerFunc(FavoritePsychePost))))
	http.Handle("/favorite/psyche", LogMiddleware(AuthMiddleware(http.HandlerFunc(FavoritePsychePost))))

	http.Handle("/test/auth/", AuthMiddleware(http.HandlerFunc(AuthTestRoute)))
	http.Handle("/test/auth", AuthMiddleware(http.HandlerFunc(AuthTestRoute)))
}

func Serve() {
	ConnectHandlers()
	log.Printf("Server running at %s...\n", PORT)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", PORT), nil); err != nil {
		log.Fatalf("Error occurred while starting http server: %v", err)
	}
}

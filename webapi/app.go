package webapi

import (
	"log"
	"net/http"
)

func Serve() {
	ConnectHandlers()
	if err := http.ListenAndServe(":61666", nil); err != nil {
		log.Fatalf("Error occurred while starting http server: %v", err)
	}
}

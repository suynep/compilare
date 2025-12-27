package web

import (
	"fmt"
	"log"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, "")
}

func getAllAeonPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w)
}

func RunServer() {
	http.HandleFunc("/", rootHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

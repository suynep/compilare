package tests

import (
	"fmt"

	"github.com/suynep/compilare/api"
)

func TestFetchBestStories() {
	api.FetchBestStories() // for the time being, this suffices
}

func TestJsonFetchBestStories() {
	fmt.Printf("Testing Best Stories Mecha\n\n")
	api.GetJsonFromPosts(api.FetchBestStories())
}

func TestJsonFetchNewStories() {
	fmt.Printf("Testing New Stories Mecha\n\n")
	api.GetJsonFromPosts(api.FetchNewStories())
}

func TestJsonFetchTopStories() {
	fmt.Printf("Testing Top Stories Mecha\n\n")
	api.GetJsonFromPosts(api.FetchTopStories())
}

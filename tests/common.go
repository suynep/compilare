package tests

import (
	"fmt"

	"github.com/suynep/compilare/api"
	"github.com/suynep/compilare/database"
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

func TestDatabaseSaves() {
	fmt.Printf("Testing Best Stories Mecha\n\n")
	data := api.GetJsonFromPosts(api.FetchBestStories())

	for _, v := range data {
		fmt.Printf("Saving %s\n", v.Title)
		database.InsertPost(v)
	}
}

func TestReadForMemoization() {
	database.ReadForMemoization()
}

package tests

import (
	"fmt"

	"github.com/suynep/compilare/api"
	"github.com/suynep/compilare/database"
	"github.com/suynep/compilare/webapi"
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

func TestBestStoriesDatabaseSaves() {
	fmt.Printf("\n\nTesting Best Stories Save Mecha\n\n")
	data := api.GetJsonFromPosts(api.FetchBestStories())

	for _, v := range data {
		fmt.Printf("Saving %s\n", v.Title)
		database.InsertPost(v, "b")
	}
}

func TestNewStoriesDatabaseSaves() {
	fmt.Printf("\n\nTesting New Stories Save Mecha\n\n")
	data := api.GetJsonFromPosts(api.FetchNewStories())

	for _, v := range data {
		fmt.Printf("Saving %s\n", v.Title)
		database.InsertPost(v, "n")
	}
}

func TestTopStoriesDatabaseSaves() {
	fmt.Printf("\n\nTesting Top Stories Save Mecha\n\n")
	data := api.GetJsonFromPosts(api.FetchTopStories())

	for _, v := range data {
		fmt.Printf("Saving %s\n", v.Title)
		database.InsertPost(v, "t")
	}
}

func TestReadForMemoization() {
	database.ReadForMemoization("t")
}

func TestWebApiServer() {
	webapi.Serve()
func TestFetchRSSFile() {
	api.FetchRSSFile()
}

func TestParseRSSFile() {
	fp := "./aeon_feed.rss"
	api.ParseRSSFile(fp)
}

func TestSaveRSSPosts() {
	fp := "./aeon_feed.rss"
	api.SaveRSSPosts(fp)
}

func TestFullFlowRSS() {
	api.FullFlowRSS()
}

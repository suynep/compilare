package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"slices"

	"github.com/suynep/compilare/database"
	"github.com/suynep/compilare/types"
)

const (
	HN_BASE_URL           = "https://hacker-news.firebaseio.com/v0/"
	HN_ROUTE_MAX_ITEM_ID  = "maxitem.json"
	HN_ROUTE_TOP_STORIES  = "topstories.json"
	HN_ROUTE_BEST_STORIES = "beststories.json"
	HN_ROUTE_NEW_STORIES  = "newstories.json"
	HN_ROUTE_ITEM_PREFIX  = "item/"
)

const TEST_LIMIT = 30

func ParseStoriesBody(body []byte) []int {
	returnList := new([]int)
	json.Unmarshal(body, returnList)
	return *returnList
}

func FetchBestStories() []int {
	url, err := url.JoinPath(HN_BASE_URL, HN_ROUTE_BEST_STORIES)
	Check(err)

	resp, err := http.Get(url)
	Check(err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	Check(err)

	initial := database.ReadForMemoization("b")
	ids := ParseStoriesBody(body)

	newIds := make([]int, 0)
	for i := range ids {
		if !slices.ContainsFunc(initial, func(e types.WebPost) bool {
			return ids[i] == e.Id
		}) {
			newIds = append(newIds, ids[i])
		}
	}

	return newIds
}

func FetchTopStories() []int {
	url, err := url.JoinPath(HN_BASE_URL, HN_ROUTE_TOP_STORIES)
	Check(err)

	resp, err := http.Get(url)
	Check(err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	Check(err)

	initial := database.ReadForMemoization("t")
	ids := ParseStoriesBody(body)

	newIds := make([]int, 0)
	for i := range ids {
		if !slices.ContainsFunc(initial, func(e types.WebPost) bool {
			return ids[i] == e.Id
		}) {
			newIds = append(newIds, ids[i])
		}
	}

	return newIds
}

func FetchNewStories() []int {
	url, err := url.JoinPath(HN_BASE_URL, HN_ROUTE_NEW_STORIES)
	Check(err)

	resp, err := http.Get(url)
	Check(err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	Check(err)

	initial := database.ReadForMemoization("n")
	ids := ParseStoriesBody(body)

	newIds := make([]int, 0)
	for i := range ids {
		if !slices.ContainsFunc(initial, func(e types.WebPost) bool {
			return ids[i] == e.Id
		}) {
			newIds = append(newIds, ids[i])
		}
	}

	return newIds
}

func GetJsonFromPosts(ids []int) []types.HNResponse {
	fetched := make([]types.HNResponse, 0)

	incompleteUrl, err := url.JoinPath(HN_BASE_URL, HN_ROUTE_ITEM_PREFIX)
	Check(err)

	ManageEmptyUrls := func(res *types.HNResponse) {
		if (*res).Url == "" {
			(*res).Url = fmt.Sprintf("https://news.ycombinator.com/item?id=%d", (*res).Id)
		}
	}

	for _, id := range ids {
		completeUrl, err := url.JoinPath(incompleteUrl, fmt.Sprintf("%d.json", id))
		Check(err)
		resp, err := http.Get(completeUrl)
		Check(err)
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		Check(err)

		// fmt.Printf("%v", string(body))
		hnResponse := new(types.HNResponse)
		json.Unmarshal(body, hnResponse)
		ManageEmptyUrls(hnResponse)
		fetched = append(fetched, *hnResponse)

		// TESTING PURPOSES
		if len(fetched) >= TEST_LIMIT {
			return fetched
		}
	}

	return fetched
}

/*
	Funcs to save fetched data to the db
*/

func SaveBestStoriesDatabase() {
	fmt.Printf("\n\nTesting Best Stories Save Mecha\n\n")
	data := GetJsonFromPosts(FetchBestStories())

	for _, v := range data {
		fmt.Printf("Saving %s\n", v.Title)
		database.InsertPost(v, "b")
	}
}

func SaveNewStoriesDatabase() {
	fmt.Printf("\n\nTesting New Stories Save Mecha\n\n")
	data := GetJsonFromPosts(FetchNewStories())

	for _, v := range data {
		fmt.Printf("Saving %s\n", v.Title)
		database.InsertPost(v, "n")
	}
}

func SaveTopStoriesDatabase() {
	fmt.Printf("\n\nTesting Top Stories Save Mecha\n\n")
	data := GetJsonFromPosts(FetchTopStories())

	for _, v := range data {
		fmt.Printf("Saving %s\n", v.Title)
		database.InsertPost(v, "t")
	}
}

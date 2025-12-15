package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

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

	return ParseStoriesBody(body)
}

func FetchTopStories() []int {
	url, err := url.JoinPath(HN_BASE_URL, HN_ROUTE_TOP_STORIES)
	Check(err)

	resp, err := http.Get(url)
	Check(err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	Check(err)

	return ParseStoriesBody(body)
}

func FetchNewStories() []int {
	url, err := url.JoinPath(HN_BASE_URL, HN_ROUTE_NEW_STORIES)
	Check(err)

	resp, err := http.Get(url)
	Check(err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	Check(err)

	return ParseStoriesBody(body)
}

func GetJsonFromPosts(ids []int) {
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
		// fmt.Printf("Extracted %d\n%d %s\n%s\n%d\n\n", id, hnResponse.Id, hnResponse.Type, hnResponse.Title, hnResponse.Score)
		fmt.Printf("%#v", hnResponse)
	}
}

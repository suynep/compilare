package api

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	HN_BASE_URL           = "https://hacker-news.firebaseio.com/v0/"
	HN_ROUTE_MAX_ITEM_ID  = "maxitem.json"
	HN_ROUTE_TOP_STORIES  = "topstories.json"
	HN_ROUTE_BEST_STORIES = "beststories.json"
	HN_ROUTE_NEW_STORIES  = "newstories.json"
	HN_ROUTE_ITEM_PREFIX  = "item/"
)

func FetchBestStories() {
	url, err := url.JoinPath(HN_BASE_URL, HN_ROUTE_BEST_STORIES)
	Check(err)

	resp, err := http.Get(url)
	Check(err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	Check(err)

	fmt.Printf("%v", string(body))
}

func FetchTopStories() {

}

func FetchNewStories() {

}

func AddPosts() {

}

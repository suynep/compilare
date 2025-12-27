package api

import (
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"os"
	"path/filepath"

	"github.com/suynep/compilare/database"
	"github.com/suynep/compilare/types"
)

const (
	AEON_BASE        = "https://aeon.co"
	AEON_ROUTE_RSS   = "/feed.rss"
	PSYCHE_BASE      = "https://psyche.co"
	PSYCHE_ROUTE_RSS = "/feed.rss"
)

var AEON_RSS_FILE_PATH = filepath.Join(".", "aeon_feed.rss")     // will be hard-coded for the time being
var PSYCHE_RSS_FILE_PATH = filepath.Join(".", "psyche_feed.rss") // will be hard-coded for the time being
var ALL_FEEDS = []string{AEON_RSS_FILE_PATH, PSYCHE_RSS_FILE_PATH}

func FetchRSSFile() {
	/*
	   For the time being, this fetches only 20 articles (titles);
	   There Should be a way of fetching more though...
	*/

	aeon_url, err := url.JoinPath(AEON_BASE, AEON_ROUTE_RSS)
	if err != nil {
		log.Fatalf("Error Occurred while fetching RSS Feed: %v", err)
	}
	psyche_url, err := url.JoinPath(PSYCHE_BASE, PSYCHE_ROUTE_RSS)
	if err != nil {
		log.Fatalf("Error Occurred while fetching RSS Feed: %v", err)
	}

	var ALL_URLS = []string{aeon_url, psyche_url}

	if len(ALL_FEEDS) == len(ALL_URLS) {
		for i, url := range ALL_URLS {

			resp, err := http.Get(url)

			if err != nil {
				log.Fatalf("Error Occurred while fetching via GET: %v", err)
			}

			defer resp.Body.Close()

			outPath, err := os.Create(ALL_FEEDS[i])

			if err != nil {
				log.Fatalf("Error Occurred while reading the response body: %v", err)
			}

			defer outPath.Close()

			_, err = io.Copy(outPath, resp.Body)

			if err != nil {
				log.Fatalf("Error Occurred while writing the response body: %v", err)
			}
		}
	} else {
		log.Fatalf("\nLength mismatch between URL slice and FEED_FILE slice\n\n")
	}
}

func ParseRSSFile(filepath string) *types.RSS {
	v := new(types.RSS)
	content, err := os.ReadFile(filepath)

	if err != nil {
		log.Fatalf("Error occurred while reading RSS file: %v", err)
	}
	err = xml.Unmarshal(content, v)
	if err != nil {
		log.Fatalf("Error occurred while unmarshalling RSS file: %v", err)
	}

	return v
}

func SaveRSSPosts(filepath string) {
	posts := ParseRSSFile(filepath)
	if strings.Contains(filepath, "psyche") {
		database.InsertPsychePosts(posts.Channel.Items)
	} else if strings.Contains(filepath, "aeon") {
		database.InsertAeonPosts(posts.Channel.Items)
	}

}

func FullFlowRSS() {
	FetchRSSFile()
	for _, fp := range ALL_FEEDS {
		SaveRSSPosts(fp)
	}

}

package database

import (
	"log"

	"github.com/suynep/compilare/types"
)

func ReadForMemoization(dataType string) []types.WebPost {
	/*
		fetching from the hnews api is expensive;
		We create the following mechanism:
			- fetch N stories every <TimeFrame> (on first run) (<TimeFrame> options: "day", "hour", "2 hours", etc. etc. -> Decide)
			- memoize the fetched data
			- Display the memoized data, instead of expensive calls every run
	*/

	q := `SELECT pid, title, url, by, score, text, time FROM posts WHERE data_type=? ORDER BY time DESC;`

	rows, err := db.Query(q, dataType)

	if err != nil {
		panic(err)
	}

	all := make([]types.WebPost, 0)

	for rows.Next() {
		current := new(types.WebPost)
		err = rows.Scan(&current.Id, &current.Title, &current.Url, &current.By, &current.Score, &current.Text, &current.Time)

		if err != nil {
			log.Fatalf("Error occurred while scanning: %v\n", err)
		}

		all = append(all, *current)
	}

	return all
}

func ReadAeonPosts() []types.Item {

	q := `SELECT title, link, creator, published, description FROM aeonposts;`

	rows, err := db.Query(q)

	if err != nil {
		panic(err)
	}

	all := make([]types.Item, 0)

	for rows.Next() {
		current := new(types.Item)
		err = rows.Scan(&current.Title, &current.Link, &current.Creator, &current.PubDate, &current.Description)

		if err != nil {
			log.Fatalf("Error occurred while scanning: %v\n", err)
		}

		all = append(all, *current)
	}

	return all
}

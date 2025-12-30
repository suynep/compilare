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

func ReadHackernewsPost(id int64) (types.WebPost, error) {
	q := `SELECT pid, title, url, by, score, text, time FROM posts WHERE id=?;`

	rows, err := db.Query(q, id)
	current := new(types.WebPost)
	err = rows.Scan(&current.Id, &current.Title, &current.Url, &current.By, &current.Score, &current.Text, &current.Time)

	if err != nil {
		return *current, err
	}

	return *current, nil
}

func ReadAeonPosts() []types.Item {
	q := `SELECT id, title, link, creator, published, description FROM aeonposts;`

	rows, err := db.Query(q)

	if err != nil {
		panic(err)
	}

	all := make([]types.Item, 0)

	for rows.Next() {
		current := new(types.Item)
		err = rows.Scan(&current.Id, &current.Title, &current.Link, &current.Creator, &current.PubDate, &current.Description)

		if err != nil {
			log.Fatalf("Error occurred while scanning: %v\n", err)
		}

		all = append(all, *current)
	}

	return all
}

func ReadAeonPost(id int64) (types.Item, error) {
	q := `SELECT id, title, link, creator, published, description FROM aeonposts WHERE id=?;`

	row := db.QueryRow(q, id)
	post := new(types.Item)

	err := row.Scan(&post.Id, &post.Title, &post.Link, &post.Creator, &post.PubDate, &post.Description)

	if err != nil {
		return *post, err
	}

	return *post, nil
}

func ReadPsychePosts() []types.Item {

	q := `SELECT id, title, link, creator, published, description FROM psycheposts;`

	rows, err := db.Query(q)

	if err != nil {
		panic(err)
	}

	all := make([]types.Item, 0)

	for rows.Next() {
		current := new(types.Item)
		err = rows.Scan(&current.Id, &current.Title, &current.Link, &current.Creator, &current.PubDate, &current.Description)

		if err != nil {
			log.Fatalf("Error occurred while scanning: %v\n", err)
		}

		all = append(all, *current)
	}

	return all
}

func ReadPsychePost(id int64) (types.Item, error) {
	q := `SELECT id, title, link, creator, published, description FROM psycheposts WHERE id=?;`

	row := db.QueryRow(q, id)
	post := new(types.Item)

	err := row.Scan(&post.Id, &post.Title, &post.Link, &post.Creator, &post.PubDate, &post.Description)

	if err != nil {
		return *post, err
	}

	return *post, nil
}

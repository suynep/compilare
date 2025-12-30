package database

import (
	"log"

	"github.com/suynep/compilare/types"
)

func InsertPost(post types.HNResponse, dataType string) {
	/*
		How... ugly...  :(
	*/
	q := `INSERT OR IGNORE INTO posts (pid,deleted,type,by,time,text,dead,parent,poll,url,score,title,parts,descendants, data_type) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`

	_, err := db.Exec(q,
		post.Id,
		post.Deleted,
		post.Type,
		post.By,
		post.Time,
		post.Text,
		post.Dead,
		post.Parent,
		post.Poll,
		// post.Kids,
		post.Url,
		post.Score,
		post.Title,
		post.Parts,
		post.Descendants,
		dataType)

	if err != nil {
		panic(err)
	}
}

func InsertPosts(limit int, posts []types.HNResponse, dataType string) {
	for i, post := range posts {
		if i >= limit {
			break
		}
		InsertPost(post, dataType)
	}
}

func InsertAeonPost(post types.Item) {
	/*
		title TEXT,
		link  TEXT,
		creator TEXT,
		published TEXT,
		description TEXT
	*/
	q := `INSERT OR IGNORE INTO aeonposts (title,link,creator,published,description) VALUES (?, ?, ?, ?, ?);`

	_, err := db.Exec(q, post.Title, post.Link, post.Creator, post.PubDate, post.Description)

	if err != nil {
		panic(err)
	}
}

func InsertPsychePost(post types.Item) {
	/*
		title TEXT,
		link  TEXT,
		creator TEXT,
		published TEXT,
		description TEXT
	*/
	q := `INSERT OR IGNORE INTO psycheposts (title,link,creator,published,description) VALUES (?, ?, ?, ?, ?);`

	_, err := db.Exec(q, post.Title, post.Link, post.Creator, post.PubDate, post.Description)

	if err != nil {
		panic(err)
	}
}

func InsertAeonPosts(posts []types.Item) {
	for _, post := range posts {
		InsertAeonPost(post)
	}
}

func InsertPsychePosts(posts []types.Item) {
	for _, post := range posts {
		InsertPsychePost(post)
	}
}

func FavoriteHackernewsPost(postId int64, user types.RegisterUser) error {
	q := `INSERT OR IGNORE INTO hackernewsfavorites (u_id, p_id) VALUES (?, ?)`

	_, err := db.Exec(q, user.Id, postId)

	if err != nil {
		log.Printf("Couldn't insert into hackernews favorites table: %v", err)
		return err
	}

	log.Printf("Favorites added successfully: hn post %d for user %d", postId, user.Id)
	return nil
}

func FavoriteAeonPost(postId int64, user types.RegisterUser) error {
	q := `INSERT OR IGNORE INTO aeonfavorites (u_id, p_id) VALUES (?, ?)`

	_, err := db.Exec(q, user.Id, postId)

	if err != nil {
		log.Printf("Couldn't insert into aeon favorites table: %v", err)
		return err
	}

	log.Printf("Favorites added successfully: aeon post %d for user %d", postId, user.Id)
	return nil
}

func FavoritePsychePost(postId int64, user types.RegisterUser) error {
	q := `INSERT OR IGNORE INTO psychefavorites (u_id, p_id) VALUES (?, ?)`

	_, err := db.Exec(q, user.Id, postId)

	if err != nil {
		log.Printf("Couldn't insert into psyche favorites table: %v", err)
		return err
	}

	log.Printf("Favorites added successfully: psyche post %d for user %d", postId, user.Id)
	return nil
}

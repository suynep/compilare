package database

import "github.com/suynep/compilare/types"

func InsertPost(post types.HNResponse, dataType string) {
	/*
		How... ugly...  :(
	*/
	q := `INSERT OR IGNORE INTO posts (id,deleted,type,by,time,text,dead,parent,poll,url,score,title,parts,descendants, data_type) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`

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

package database

import "github.com/suynep/compilare/types"

func InsertPost(post types.HNResponse) {
	/*
		How... ugly...  :(
	*/
	q := `INSERT OR IGNORE INTO posts (id,deleted,type,by,time,text,dead,parent,poll,url,score,title,parts,descendants) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`

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
		post.Descendants)

	if err != nil {
		panic(err)
	}
}

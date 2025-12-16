package types

import "time"

type HNResponse struct {
	Id          int    `json:"id,omitempty" db:"id,omitempty"`
	Deleted     bool   `json:"deleted,omitempty" db:"deleted,omitempty"`
	Type        string `json:"type,omitempty" db:"type,omitempty"`
	By          string `json:"by,omitempty" db:"by,omitempty"`
	Time        int64  `json:"time,omitempty" db:"time,omitempty"`
	Text        string `json:"text,omitempty" db:"text,omitempty"`
	Dead        bool   `json:"dead,omitempty" db:"dead,omitempty"`
	Parent      int    `json:"parent,omitempty" db:"parent,omitempty"`
	Poll        string `json:"poll,omitempty" db:"poll,omitempty"`
	Kids        []int  `json:"kids,omitempty" db:"kids,omitempty"`
	Url         string `json:"url,omitempty" db:"url,omitempty"`
	Score       int    `json:"score,omitempty" db:"score,omitempty"`
	Title       string `json:"title,omitempty" db:"title,omitempty"`
	Parts       string `json:"parts,omitempty" db:"parts,omitempty"`
	Descendants int    `json:"descendants,omitempty" db:"descendants,omitempty"`
}

type Config struct {
	Run struct {
		Time time.Time `json:"time,omitempty"`
	} `json:"run,omitempty"`
}

type WebPost struct {
	/*
		type for displaying the actual data to the web
	*/
	Id    int    `json:"id,omitempty" db:"id,omitempty"`
	By    string `json:"by,omitempty" db:"by,omitempty"`
	Url   string `json:"url,omitempty" db:"url,omitempty"`
	Score int    `json:"score,omitempty" db:"score,omitempty"`
	Title string `json:"title,omitempty" db:"title,omitempty"`
	Time  int64  `json:"time,omitempty" db:"time,omitempty"`
	Text  string `json:"text,omitempty" db:"text,omitempty"`
}

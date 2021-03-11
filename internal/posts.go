package internal

var PostCounter int = 0
var Posts = make(map[int]Post)

type Post struct {
	Id         int `json:"id"`
	AuthorId   string
	AuthorName string
	Text       string
	UrlImg     string
}


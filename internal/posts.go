package internal

var PostCounter int = 0
var Posts = make(map[int]Post)

type Post struct {
	Id         int
	AuthorId   string
	AuthorName string `json:"postCreator"`
	Text       string `json:"textPost"`
	UrlImg     string `json:"imgContent"`
}

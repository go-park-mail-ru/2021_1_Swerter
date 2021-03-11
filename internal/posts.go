package internal

type Post struct {
	Id         int `json:"id"`
	AuthorId   string
	AuthorName string
	Text       string
	UrlImg     string
}

func NewPost(id int, authorId string, authorName string, text string, urlImg string) Post {
	return Post{
		Id:         id,
		AuthorId:   authorId,
		AuthorName: authorName,
		Text:       text,
		UrlImg:     urlImg,
	}
}

var Posts = make(map[int]Post)

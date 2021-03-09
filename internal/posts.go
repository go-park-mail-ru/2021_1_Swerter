package internal

type Post struct {
	Id         int `json:"id"`
	AuthorName string
	Text       string
	UrlImg     string
}

func NewPost(id int, authorName string, text string, urlImg string) Post {
	return Post{
		Id: id,
		AuthorName: authorName,
		Text: text,
		UrlImg: urlImg,
	}
}

var Posts = make(map[int]Post)





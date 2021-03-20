package internal

var PostCounter int = 0
var Posts = make(map[int]Post)

type Post struct {
	Id        int
	Author    string `json:"postCreator"`//устанавливаю при подзгрузке
	AuthorAva string `json:"imgAvatar"` //устанавливаю при подзгрузке
	AuthorId  string `json:"postCreatorId"`
	Text      string `json:"textPost"`
	UrlImg    string `json:"imgContent"`
	Date      string `json:"date"`
}

package internal

var IDToLogin = make(map[string]string)

//Login: User
var Users = map[string]User{}
var IDCounter int

type User struct {
	ID        string       `json:"userId"`
	Login     string       `json:"login"`
	FirstName string       `json:"firstName"`
	LastName  string       `json:"lastName"`
	Password  string       `json:"-"`
	Posts     map[int]Post `json:"postsData"`
	Avatar    string       `json:"avatar"`
}
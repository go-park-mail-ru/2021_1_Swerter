package internal

var IDToLogin map[string]string = make(map[string]string)
//Login: User
var Users map[string]User = map[string]User{}
var IDCounter int

type User struct {
	ID        string
	Login     string
	FirstName string
	LastName  string
	Password  string
}

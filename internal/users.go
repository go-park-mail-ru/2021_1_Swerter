package internal

//key - User.login
var Users map[string]User = map[string]User{}
var IDCounter int

type User struct {
	ID        string
	Login     string
	FirstName string
	LastName  string
	Password  string
}

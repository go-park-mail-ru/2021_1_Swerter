package internal

var Users map[string]User = map[string]User{}
var IDCounter int

type User struct {
	ID        int
	Login     string
	FirstName string
	LastName  string
	Password  string
}

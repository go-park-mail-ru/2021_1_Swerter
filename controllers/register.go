package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"vk.com/models"
)

var Users []models.User

var idCounter int
var registerSuccess string = `"status":"true"`
var registerFail string = `"status":"false"`

func RegisterPage(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != http.MethodPost {
		return
	}

	decoder := json.NewDecoder(r.Body)
	newUser := models.User{}
	decoder.Decode(&newUser)
	idCounter++
	newUser.ID = idCounter
	fmt.Println("Im in reg")
	Users = append(Users, newUser)
	fmt.Printf("%+v\n", newUser)
	w.Write([]byte(registerSuccess))
	http.Redirect(w, r, "/", http.StatusFound)
}

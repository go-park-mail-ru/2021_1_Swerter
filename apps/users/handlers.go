package users

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	i "my-motivation/internal"
	"my-motivation/utils"
	"net/http"
)

func userProfile(w http.ResponseWriter, r *http.Request) {
	utils.SetupCORS(&w)

	if r.Method == http.MethodOptions {
		return
	}

	if r.Method == http.MethodGet {
		getUserProfile(w, r)
	}

	if r.Method == http.MethodPost {
		updateUserProfile(w, r)
	}
}

func getUserProfile(w http.ResponseWriter, r *http.Request) {
	utils.SetupCORS(&w)
	user := utils.SessionToUser(r)
	if user == nil {
		w.WriteHeader(http.StatusForbidden)
	}
	userJson, _ := json.Marshal(user)
	w.Write(userJson)

}

func updateUserProfile(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	newUser := i.User{}
	decoder.Decode(&newUser)

	oldUser := utils.SessionToUser(r)
	if oldUser == nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	updateUser(&newUser, oldUser)
	log.Printf("User update success: %+v\n", newUser)
	return
}


func getUserProfileByID(w http.ResponseWriter, r *http.Request) {
	utils.SetupCORS(&w)
	u := i.Users[i.IDToLogin[mux.Vars(r)["userID"]]]
	log.Println("get user with id:", mux.Vars(r)["userID"])
	body, _ := json.Marshal(&u)
	w.Write(body)
}

func updateUser(newUser *i.User, oldUser *i.User) {
	newUser.ID = oldUser.ID

	if newUser.Login == "" {
		newUser.Login = oldUser.Login
	} else {
		i.IDToLogin[newUser.ID] = newUser.Login
	}

	if newUser.Password == "" {
		newUser.Password = oldUser.Password
	} else {
		newUser.Password = utils.HashPassword(newUser.Password)
	}

	if newUser.FirstName == "" {
		newUser.FirstName = oldUser.FirstName
	}

	if newUser.LastName == "" {
		newUser.LastName = oldUser.LastName
	}

	newUser.Posts = oldUser.Posts

	delete(i.Users, oldUser.Login)
	i.Users[newUser.Login] = *newUser
}
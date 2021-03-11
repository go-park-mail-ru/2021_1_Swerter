package news

import (
	"encoding/json"
	"fmt"
	"log"
	i "my-motivation/internal"
	u "my-motivation/utils"
	"net/http"
)

func allPosts(w http.ResponseWriter, r *http.Request) {
	u.SetupCORS(&w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	user := u.SessionToUser(r)
	if user == nil {
		log.Println("Need auth")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	jsonValue, _ := json.Marshal(i.Posts)
	fmt.Println(jsonValue)
	w.Write([]byte(jsonValue))
	w.WriteHeader(http.StatusOK)
}

func addPost(w http.ResponseWriter, r *http.Request) {
	u.SetupCORS(&w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	user := u.SessionToUser(r)
	if user == nil {
		log.Println("Add post failed")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	storePost(user, r)
	w.WriteHeader(http.StatusOK)
}

func storePost(user *i.User, r *http.Request) {
	i.PostCounter++
	decoder := json.NewDecoder(r.Body)
	newPost := i.Post{}
	decoder.Decode(&newPost)
	newPost.AuthorId = user.ID
	newPost.Id = i.PostCounter
	i.Posts[i.PostCounter] = newPost
	newPosts := append(i.Users[i.IDToLogin[user.ID]].Posts, newPost.Id)
	oldUser := i.Users[i.IDToLogin[user.ID]]
	oldUser.Posts = newPosts
	i.Users[i.IDToLogin[user.ID]] = oldUser
	fmt.Printf("New post. Post data: %+v\n", newPost)
}

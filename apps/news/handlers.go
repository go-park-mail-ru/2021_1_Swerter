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
	user.Posts = append(user.Posts, newPost.Id)
}
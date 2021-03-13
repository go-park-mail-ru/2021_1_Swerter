package posts

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	i "my-motivation/internal"
	u "my-motivation/utils"
	"net/http"
	"os"
	"time"
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
	newPost := i.Post{}
	newPost.Id = i.PostCounter
	newPost.AuthorId = user.ID
	newPost.Date = r.FormValue("date")
	newPost.Text = r.FormValue("textPost")
	storeImg(r, &newPost)
	i.Posts[i.PostCounter] = newPost
	i.Users[i.IDToLogin[user.ID]].Posts[newPost.Id] = newPost
	fmt.Printf("New post. Post data: %+v\n", newPost)
}

func storeImg(r *http.Request, post *i.Post) {
	imgContent, handler, err := r.FormFile("imgContent")
	if err != nil {
		fmt.Printf("No post img content\n")
	}

	t := time.Now()
	salt := fmt.Sprintf(t.Format(time.RFC3339))
	fileName := u.Hash(handler.Filename + salt)

	defer imgContent.Close()
	localImg, err := os.OpenFile("./static/posts/" + fileName, os.O_WRONLY|os.O_CREATE, 0666)
	post.UrlImg = "/static/posts/" + fileName
	if err != nil {
		fmt.Printf("Cant create file\n")
	}

	defer localImg.Close()
	_, _ = io.Copy(localImg, imgContent)
	fmt.Printf("Load new file\n")
}

package http

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"my-motivation/internal/app/models"
	"net/http"
)

type PostHandler struct {
	PostUsecase models.PostsUsecase
}

func NewPostHandler(r *mux.Router, pu models.PostsUsecase) {
	handler := &PostHandler{
		PostUsecase: pu,
	}
	r.HandleFunc("/posts", handler.allPosts).Methods("GET", "OPTIONS")
	r.HandleFunc("/posts/add", handler.addPost).Methods("POST", "OPTIONS")
}

func (ph *PostHandler) allPosts(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	session, err := r.Cookie("session_id")
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	posts, err := ph.PostUsecase.GetPosts(r.Context(), session.Value)
	if err != nil {
		return
	}
	jsonValue, _ := json.Marshal(posts)
	fmt.Println(jsonValue)
	w.Write([]byte(jsonValue))
	w.WriteHeader(http.StatusOK)
}

func (ph *PostHandler) addPost(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	defer r.Body.Close()
	session, err := r.Cookie("session_id")
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	imgFile, fileHandler, err := r.FormFile("imgContent")
	newPost := models.Post{}
	newPost.Date = r.FormValue("date")
	newPost.Text = r.FormValue("textPost")

	err = ph.PostUsecase.SavePost(r.Context(), session.Value, imgFile, fileHandler, &newPost)

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	w.WriteHeader(http.StatusOK)
}
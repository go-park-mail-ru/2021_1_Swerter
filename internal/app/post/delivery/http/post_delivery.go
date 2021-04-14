package http

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"my-motivation/internal/app/models"
	"my-motivation/internal/pkg/utils/logger"
	"net/http"
)

type PostHandler struct {
	PostUsecase models.PostsUsecase
	logger logger.LoggerModel
}

func NewPostHandler(r *mux.Router, pu models.PostsUsecase, l logger.LoggerModel) {
	handler := &PostHandler{
		PostUsecase: pu,
		logger: l,
	}
	r.HandleFunc("/posts", handler.allPosts).Methods("GET", "OPTIONS")
	r.HandleFunc("/posts/add", handler.addPost).Methods("POST", "OPTIONS")
}

func (ph *PostHandler) allPosts(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	session, err := r.Cookie("session_id")
	if err != nil {
		ph.logger.Error("no authorization")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	posts, err := ph.PostUsecase.GetPosts(r.Context(), session.Value)
	if err != nil {
		ph.logger.Error(err.Error())
		return
	}
	jsonValue, _ := json.Marshal(posts)
	ph.logger.Debug("Posts sends")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jsonValue))
}


func (ph *PostHandler) addPost(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	session, err := r.Cookie("session_id")
	if err != nil {
		ph.logger.Error("no authorization")
		w.WriteHeader(http.StatusForbidden)
		return
	}


	mr, err := r.MultipartReader()
	form, err := mr.ReadForm(100000)
	if err != nil {
		ph.logger.Error("no files")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	newPost := models.Post{}
	newPost.Date = form.Value["date"][0]
	newPost.Text = form.Value["textPost"][0]

	err = ph.PostUsecase.SavePost(r.Context(), session.Value, form.File, &newPost)

	if err != nil {
		ph.logger.Error(err.Error())
		w.WriteHeader(http.StatusForbidden)
		return
	}

	w.WriteHeader(http.StatusOK)
}
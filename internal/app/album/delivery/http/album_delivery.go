package http

import (
	"github.com/gorilla/mux"
	"my-motivation/internal/app/models"
	"my-motivation/internal/pkg/utils/logger"
	"net/http"
)

type AlbumHandler struct {
	AlbumUsecase models.AlbumsUsecase
	logger logger.LoggerModel
}

func NewAlbumHandler(r *mux.Router, au models.AlbumsUsecase, l logger.LoggerModel) {
	handler := &AlbumHandler{
		AlbumUsecase: au,
		logger: l,
	}
	//r.HandleFunc("/posts", handler.allPosts).Methods("GET", "OPTIONS")
	r.HandleFunc("/album/add", handler.addAlbum).Methods("POST", "OPTIONS")
}

func (ah *AlbumHandler) addAlbum(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	session, err := r.Cookie("session_id")
	if err != nil {
		ah.logger.Error("no authorization")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	mr, err := r.MultipartReader()
	//Можно ли так
	form, err := mr.ReadForm(100000)
	if err != nil {
		ah.logger.Error("no files")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	newAlbum := models.Album{}
	newAlbum.Title = form.Value["albumTitle"][0]
	newAlbum.Description = form.Value["albumDescription"][0]

	err = ah.AlbumUsecase.SaveAlbum(r.Context(), session.Value, form.File, &newAlbum)

	if err != nil {
		ah.logger.Error(err.Error())
		w.WriteHeader(http.StatusForbidden)
		return
	}

	w.WriteHeader(http.StatusOK)
}
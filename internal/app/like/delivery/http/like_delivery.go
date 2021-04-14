package http

import (
	"github.com/gorilla/mux"
	"my-motivation/internal/app/models"
	"my-motivation/internal/pkg/utils/logger"
	"net/http"
	"strconv"
)

type LikeHandler struct {
	LikeUsecase models.LikeUsecase
	logger      logger.LoggerModel
}

func NewLikeHandler(r *mux.Router, lu models.LikeUsecase, l logger.LoggerModel) {
	handler := &LikeHandler{
		LikeUsecase: lu,
		logger:      l,
	}
	r.HandleFunc("/like/post/{postID}", handler.changeLike).Methods("POST", "OPTIONS")
}

func (lh *LikeHandler) changeLike(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if r.Method == http.MethodOptions {
		return
	}

	session, err := r.Cookie("session_id")
	if err != nil || session == nil {
		lh.logger.Error("no authorization")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	ctx := r.Context()
	postId, err := strconv.Atoi(mux.Vars(r)["postID"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = lh.LikeUsecase.ChangeLike(ctx, session.Value, postId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
}

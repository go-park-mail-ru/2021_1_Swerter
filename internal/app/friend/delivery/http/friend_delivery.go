package http

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"my-motivation/internal/app/models"
	"my-motivation/internal/pkg/utils/logger"
	"net/http"
)

type FriendHandler struct {
	FriendUsecase models.FriendUsecase
	logger      logger.LoggerModel
}

func NewFiendHandler(r *mux.Router, fu models.FriendUsecase, l *logger.Logger) {
	handler := &FriendHandler{
		FriendUsecase: fu,
		logger:      l,
	}

	r.HandleFunc("/user/friend/add", handler.addFriend).Methods("POST", "OPTIONS")
	r.HandleFunc("/user/friends", handler.getFriends).Methods("GET", "OPTIONS")
}

func (fh *FriendHandler) getFriends(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if r.Method == http.MethodOptions {
		return
	}

	session, err := r.Cookie("session_id")
	if err != nil || session == nil{
		fh.logger.Error("no authorization")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	ctx := r.Context()
	users, err := fh.FriendUsecase.GetFriends(ctx, session.Value)
	if err != nil {
		fh.logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//TODO: не отправляь поля pass, login
	jsonUsers, err := json.Marshal(users)
	if err != nil {
		fh.logger.Error("can`t marshal friends")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fh.logger.Debug("send all friends")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonUsers)
}

func (fh *FriendHandler) addFriend(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if r.Method == http.MethodOptions {
		return
	}

	session, err := r.Cookie("session_id")
	if err != nil || session == nil{
		fh.logger.Error("no authorization")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	decoder := json.NewDecoder(r.Body)
	user := &models.User{}
	err = decoder.Decode(user)
	if err != nil {
		fh.logger.Error(err.Error())
		w.WriteHeader(http.StatusNoContent)
		return
	}

	ctx := r.Context()
	err = fh.FriendUsecase.AddFriend(ctx, session.Value, user)
	if err != nil {
		fh.logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fh.logger.Debug("Add new friend")
	w.WriteHeader(http.StatusOK)
}

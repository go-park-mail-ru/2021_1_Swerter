package http

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"my-motivation/internal/app/models"
	"my-motivation/internal/pkg/utils/logger"
	"net/http"
	"strconv"
	"time"
)

type UserHandler struct {
	UserUsecase models.UserUsecase
	logger      logger.LoggerModel
}

func NewUserHandler(r *mux.Router, uu models.UserUsecase, l *logger.Logger) {
	handler := &UserHandler{
		UserUsecase: uu,
		logger:      l,
	}

	r.HandleFunc("/profile/loadImg", handler.uploadAvatar).Methods("POST", "OPTIONS")
	r.HandleFunc("/profile", handler.userProfile).Methods("GET", "POST", "OPTIONS")
	r.HandleFunc("/profile/{userID}", handler.getUserProfileByID).Methods("GET", "OPTIONS")

	r.HandleFunc("/login", handler.login).Methods("POST", "OPTIONS")
	r.HandleFunc("/logout", handler.logout).Methods("POST", "OPTIONS")
	r.HandleFunc("/register", handler.register).Methods("POST", "OPTIONS")
}

func (uh *UserHandler) uploadAvatar(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if r.Method == http.MethodOptions {
		return
	}

	file, handler, err := r.FormFile("avatar")
	defer file.Close()

	uh.logger.Debug(fmt.Sprintf("Upload avatar: %s ", handler.Header))
	if err != nil {
		uh.logger.Error(err.Error())
		return
	}

	session, err := r.Cookie("session_id")
	if err != nil || session == nil {
		uh.logger.Error("no authorization")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	err = uh.UserUsecase.UploadAvatar(r.Context(), session.Value, file)
	if err != nil {
		uh.logger.Error(err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (uh *UserHandler) getUserProfileByID(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if r.Method == http.MethodOptions {
		return
	}

	userId, err := strconv.Atoi(mux.Vars(r)["userID"][2:])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	user, err := uh.UserUsecase.GetUserById(r.Context(), userId)
	if err != nil {
		uh.logger.Error(err.Error())
		w.WriteHeader(http.StatusNotFound)
		return
	}

	uh.logger.Debug(fmt.Sprintf("get user with id: %d", userId))
	body, _ := json.Marshal(user)
	w.Write(body)
}

func (uh *UserHandler) userProfile(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method == http.MethodOptions {
		return
	}

	if r.Method == http.MethodGet {
		session, err := r.Cookie("session_id")
		if err != nil || session == nil {
			uh.logger.Error("no authorization")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		user, err := uh.UserUsecase.GetUserBySession(r.Context(), session.Value)
		if err != nil || user == nil {
			uh.logger.Error(err.Error())
			w.WriteHeader(http.StatusUnauthorized)
		}
		userJson, _ := json.Marshal(user)
		w.Write(userJson)
	}

	if r.Method == http.MethodPost {
		session, err := r.Cookie("session_id")
		if err != nil {
			uh.logger.Error("no authorization")
			w.WriteHeader(http.StatusForbidden)
			return
		}
		decoder := json.NewDecoder(r.Body)
		newUser := &models.User{}
		err = decoder.Decode(newUser)
		if err != nil {
			return
		}

		err = uh.UserUsecase.UpdateUser(r.Context(), newUser, session.Value)
		if err != nil {
			uh.logger.Error(err.Error())
			w.WriteHeader(http.StatusForbidden)
			return
		}
		w.WriteHeader(http.StatusOK)
		uh.logger.Debug(fmt.Sprintf("User update success: %+v\n", newUser))
		return
	}
}

func (uh *UserHandler) login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	user := models.User{}
	err := decoder.Decode(&user)
	if err != nil {
		uh.logger.Error(err.Error())
		return
	}

	ctx := r.Context()
	session, err := uh.UserUsecase.LoginUser(ctx, &user)
	if err != nil {
		uh.logger.Error("no authorization")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	uh.logger.Debug(fmt.Sprintf("New session: %+v\n", session))
	expiration := time.Now().Add(10 * time.Hour)
	cookie := http.Cookie{
		Name:    "session_id",
		Value:   session.ID,
		Expires: expiration,
		SameSite: http.SameSiteNoneMode,
		Secure: true,
	}

	http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusOK)
}

func (uh *UserHandler) logout(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	session, err := r.Cookie("session_id")

	if err == http.ErrNoCookie {
		uh.logger.Error("No cookie was provided for logout")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	session.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, session)
}

func (uh *UserHandler) register(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	newUser := models.User{}
	decoder.Decode(&newUser)

	ctx := r.Context()
	err := uh.UserUsecase.SaveUser(ctx, &newUser)
	if err != nil {
		uh.logger.Error(err.Error())
		w.WriteHeader(http.StatusForbidden)
		return
	}

	uh.logger.Debug(fmt.Sprintf("New user. Private user data: %+v\n", newUser))

	responseBody := []byte("{\"userID\":" + fmt.Sprint(newUser.ID) + "}")
	w.Write(responseBody)
}

package http

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"my-motivation/internal/app/models"
	"net/http"
	"time"
)

type UserHandler struct {
	UserUsecase models.UserUsecase
}

func NewUserHandler(r *mux.Router, uu models.UserUsecase) {
	handler := &UserHandler{
		UserUsecase: uu,
	}
	//user
	r.HandleFunc("/profile/loadImg", handler.uploadAvatar).Methods("POST")
	r.HandleFunc("/profile", handler.userProfile).Methods("GET", "POST", "OPTIONS")
	r.HandleFunc("/profile/{userID}", handler.getUserProfileByID).Methods("GET", "OPTIONS")
	//auth
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

	fmt.Println(handler.Header, err)
	if err != nil {
		log.Fatal(err)
		return
	}
	session, err := r.Cookie("session_id")
	if err != nil || session == nil{
		w.WriteHeader(http.StatusForbidden)
		return
	}
	err = uh.UserUsecase.UploadAvatar(r.Context(), session.Value, file)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (uh *UserHandler) getUserProfileByID(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if r.Method == http.MethodOptions {
		return
	}
	user, err := uh.UserUsecase.GetUserById(r.Context(), mux.Vars(r)["userID"])
	if err != nil {
		log.Printf("no user with %s id\n", mux.Vars(r)["userID"])
		w.WriteHeader(http.StatusNotFound)
	}
	log.Println("get user with id:", mux.Vars(r)["userID"])
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
		if err != nil || session == nil{
			w.WriteHeader(http.StatusForbidden)
			return
		}
		user, err := uh.UserUsecase.GetUserBySession(r.Context(), session.Value)
		if err != nil || user == nil {
			w.WriteHeader(http.StatusForbidden)
		}
		userJson, _ := json.Marshal(user)
		w.Write(userJson)
	}

	if r.Method == http.MethodPost {
		session, err := r.Cookie("session_id")
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		decoder := json.NewDecoder(r.Body)
		newUser := &models.User{}
		err = decoder.Decode(newUser)
		if err != nil {
			return
		}

		uh.UserUsecase.UpdateUser(r.Context(), newUser, session.Value)
		w.WriteHeader(http.StatusOK)
		log.Printf("User update success: %+v\n", newUser)
		return
	}
}

func (uh *UserHandler) login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	user := models.User{}
	err := decoder.Decode(&user)
	if err != nil {
		w.Write([]byte("can not decode user"))
		return
	}

	ctx := r.Context()
	session, err := uh.UserUsecase.LoginUser(ctx, &user)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	log.Printf("New session: %+v\n", session)
	expiration := time.Now().Add(10 * time.Hour)
	cookie := http.Cookie{
		Name:    "session_id",
		Value:   session.ID,
		Expires: expiration,
		//SameSite: http.SameSiteNoneMode,
		//Secure: true,
	}

	http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusOK)
}

func (uh *UserHandler) logout(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	session, err := r.Cookie("session_id")

	if err == http.ErrNoCookie {
		log.Println("No cookie was provided for logout")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Println("Logout")

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
		w.Write([]byte(err.Error()))
		return
	}
	fmt.Printf("New user. Private user data: %+v\n", newUser)

	responseBody := []byte("{\"userID\":" + newUser.ID + "}")
	w.Write(responseBody)
}
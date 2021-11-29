package webauth

import (
	"Dp218Go/auth"
	"Dp218Go/model"
	"Dp218Go/repository"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	ErrSignUp = errors.New("signup error")
	ErrSignIn = errors.New("signin error")
)

const sessionName = "login"

type AuthService struct {
	DB        repository.UserRepo
	sessStore sessions.Store
}

func NewAuthService(db repository.UserRepo, store sessions.Store) *AuthService {
	return &AuthService{
		DB:        db,
		sessStore: store,
	}
}

func (sv *AuthService) SignUp(w http.ResponseWriter, r *http.Request) {

	// TODO implement validation
	user := model.User{
		Name:     r.FormValue("name"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	pass, err := auth.HashPassword(user.Password)
	if err != nil {
		fmt.Println(err)
		http.Error(w, ErrSignUp.Error(), http.StatusInternalServerError)
		return
	}
	user.Password = pass

	_, err = sv.DB.Create(r.Context(), user)
	if err != nil {
		fmt.Println(err)
		http.Error(w, ErrSignUp.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/login", http.StatusFound)
}

func (sv *AuthService) SignIn(w http.ResponseWriter, r *http.Request) {
	// // TODO implement validation
	type authRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	req := authRequest{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	// // check if user exists
	user, err := sv.DB.GetByEmail(r.Context(), req.Email)
	if err != nil {
		fmt.Println(err)
		http.Error(w, ErrSignIn.Error(), http.StatusForbidden)
		return
	}

	if err := auth.CheckPassword(user.Password, req.Password); err != nil {
		fmt.Println(err)
		http.Error(w, ErrSignIn.Error(), http.StatusForbidden)
		return
	}

	session, err := sv.sessStore.Get(r, sessionName)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["id"] = user.ID
	err = session.Save(r, w)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/main", http.StatusFound)
}

func (sv *AuthService) SignOut(w http.ResponseWriter, r *http.Request) {
	session, err := sv.sessStore.Get(r, sessionName)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["id"] = nil
	session.Options.MaxAge = -1
	err = session.Save(r, w)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/login", http.StatusFound)
}

func (sv *AuthService) SessMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		sess, err := sv.sessStore.Get(r, sessionName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if sess.IsNew {
			http.Error(w, "not authorized", http.StatusForbidden)
			return
		}

		next(w, r)
	}
}

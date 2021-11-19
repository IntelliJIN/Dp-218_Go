package session

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/ITA-Dnipro/Dp-218_Go/model"
)

var (
	ErrSignUp = errors.New("signup error")
	ErrSignIn = errors.New("signin error")
)

func (sv *AuthService) SignUp(w http.ResponseWriter, r *http.Request) {
	// TODO refactor
	// if r.Method == "GET" {
	// 	http.ServeFile(w, r, "./static/html/login-registration.html")
	// }

	// TODO implement validation
	user := model.User{
		Name:     r.FormValue("name"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	// // check if user already exists
	// TODO change. for map
	if u, err := sv.DB.GetByEmail(r.Context(), user.Email); err != nil {
		fmt.Println(err)
		if u != nil {
			http.Error(w, fmt.Errorf("user exists").Error(), http.StatusForbidden)
			return
		}
	}

	err := user.HashPassword()
	if err != nil {
		fmt.Println(err)
		http.Error(w, ErrSignUp.Error(), http.StatusInternalServerError)
		return
	}

	_, err = sv.DB.Create(r.Context(), user)
	if err != nil {
		fmt.Println(err)
		http.Error(w, ErrSignUp.Error(), http.StatusInternalServerError)
		return
	}
	// nuser.Sanitize()

	u, _ := sv.DB.GetByEmail(r.Context(), user.Email)
	fmt.Println("signed up user: ", u)
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
	fmt.Println(req)
	// // check if user exists
	user, err := sv.DB.GetByEmail(r.Context(), req.Email)
	if err != nil {
		fmt.Println(err)
		http.Error(w, ErrSignIn.Error(), http.StatusForbidden)
		return
	}

	if err := user.CheckPassword(req.Password); err != nil {
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

	session.Values["email"] = user.Email
	err = session.Save(r, w)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/main", http.StatusFound)
}

package main

import (
	"Dp218Go/auth/webauth"
	"Dp218Go/model"
	"Dp218Go/repository/mapdb"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

const address = "localhost:9000"

var sessKey = "secretkey"

func main() {
	repo := mapdb.NewUserRepoMap(map[string]*model.User{})
	sessStore := sessions.NewCookieStore([]byte(sessKey))

	auth := webauth.NewAuthService(repo, sessStore)

	http.HandleFunc("/signup", auth.SignUp)
	http.HandleFunc("/signin", auth.SignIn)

	fs := http.FileServer(http.Dir("../static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("server started and listening: ", address)
	log.Fatal("server stopped ", http.ListenAndServe(address, nil))
}

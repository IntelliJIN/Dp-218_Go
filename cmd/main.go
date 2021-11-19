package main

import (
	"log"
	"net/http"

	"github.com/ITA-Dnipro/Dp-218_Go/auth/session"
	"github.com/ITA-Dnipro/Dp-218_Go/model"
	"github.com/ITA-Dnipro/Dp-218_Go/repository/mapdb"
	"github.com/gorilla/sessions"
)

const address = "localhost:9000"

var sessKey = "secretkey"

func main() {
	repo := mapdb.NewUserRepoMap(map[string]*model.User{})
	sessStore := sessions.NewCookieStore([]byte(sessKey))

	auth := session.NewAuthService(repo, sessStore)

	http.HandleFunc("/signup", auth.SignUp)
	http.HandleFunc("/signin", auth.SignIn)

	fs := http.FileServer(http.Dir("../static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("server started and listening: ", address)
	log.Fatal("server stopped ", http.ListenAndServe(address, nil))
}

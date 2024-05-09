package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ilmsg/fictional-barnacle/handler"
)

func main() {
	fmt.Println("hi")

	r := mux.NewRouter()

	public := http.FileServer(http.Dir("./public"))
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", public))

	r.HandleFunc("/users/register", handler.Register)
	r.HandleFunc("/users/login", handler.Login)
	r.HandleFunc("/users/profile", handler.Profile)
	r.HandleFunc("/users/setting", handler.Setting)
	r.HandleFunc("/users/logout", handler.Logout)

	r.HandleFunc("/posts/{id}", handler.PostDetail)
	r.HandleFunc("/posts", handler.PostList)
	r.HandleFunc("/", handler.Home)

	http.ListenAndServe(":3030", r)
}

package main

import (
	"SSO-Snap/Services/server"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	http.Handle("/", r)
	r.HandleFunc("/account", login).Methods("GET")
	r.HandleFunc("/account", server.Login).Methods("POST")
	r.HandleFunc("/logout", server.LogoutHandler).Methods("GET")
	r.HandleFunc("/user/{username}", server.GetUserHandler).Methods("GET")
	fmt.Println("server is runing ....")
	http.ListenAndServe(":8081", nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./template/login.html")
}

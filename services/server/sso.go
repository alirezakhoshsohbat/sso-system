package server

import (
	Database "SSO-Snap/Services/db"
	"database/sql"
	"fmt"
	"net/http"
	// "time"

	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var (
	my_Secret_Key = []byte("Alireza9268")
	store         = sessions.NewCookieStore(my_Secret_Key)
	db, _         = Database.DatabaseConection()
)

func Login(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	if checkCredentials(username, password) {
        current_time := time.Now()
        expiration_time := current_time.Add(time.Minute * 1)
        store.Options = &sessions.Options{
            MaxAge: int(expiration_time.Sub(current_time).Seconds()),
            Path:   "/",
        }
		session, _ := store.New(r, "sso-session")
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["user"] = true
		tokenString, err := token.SignedString(my_Secret_Key)
		if err != nil {
			fmt.Println("خطا در امضای توکن:", err)
			return
		}
		session.Values["access_token"] = tokenString
		session.Save(r, w)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "login success")
		
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Login failed")
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "sso-session")
	session.Options.MaxAge = -1
	session.Save(r, w)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Logout successful")
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	session, _ := store.Get(r, "sso-session")
	storedUsername, ok := session.Values["username"].(string)
	if !ok || storedUsername != username {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Access denied")
		return
	}
	userInfo, err := getUserInfo(username)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"username": "%s", "password": "%s", "adress": "%s", "phone": "%s"}`, userInfo["username"], userInfo["password"], userInfo["adress"], userInfo["phone"])
}

func checkCredentials(username, password string) bool {
	query := "SELECT username, password FROM users.login WHERE username=? AND password=?;"
	var dbUsername, dbPassword string
	err := db.QueryRow(query, username, password).Scan(&dbUsername, &dbPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		} else {
			fmt.Println("query error: ", err)
			return false
		}
	}
	return true
}

func getUserInfo(username string) (map[string]string, error) {
	query := "SELECT username, password, adress, phone FROM users.login WHERE username=? ;"

	var dbUsername, dbPassword, dbAddress, dbPhone string

	err := db.QueryRow(query, username).Scan(&dbUsername, &dbPassword, &dbAddress, &dbPhone)
	if err != nil {
		return nil, err
	}

	userinfo := make(map[string]string)
	userinfo["username"] = dbUsername
	userinfo["password"] = dbPassword
	userinfo["address"] = dbAddress
	userinfo["phone"] = dbPhone

	return userinfo, nil
}

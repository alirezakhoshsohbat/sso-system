package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/sessions"
)

var (
	secretKey = []byte("Alireza9268")
	
	store = sessions.NewCookieStore(secretKey)
	authServerURL = "http://localhost:8081"
)

func main() {
	http.HandleFunc("/t", protectedResourceHandler)
	fmt.Println("client 2 is runing ....")
	http.ListenAndServe(":8083", nil)
}

func protectedResourceHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "sso-session")
	accessToken := session.Values["access_token"]
	if accessToken == nil {
		http.Redirect(w, r, authServerURL+"/login", http.StatusFound)
		return
	}
	fmt.Fprintf(w, "Protected resource accessed successfully!")
}






// 
// 
// package main
// 
// import (
//     "fmt"
//     "net/http"
//     "github.com/dgrijalva/jwt-go"
//     "github.com/dgrijalva/jwt-go/request"
// )
// 
// func main() {
//     // کلید مخفی برای تایید امضا JWT
//     secretKey := []byte("your-secret-key")
// 
//     // درخواست HTTP که حاوی JWT است
//     // در اینجا می‌توانید JWT را از درخواست HTTP دریافت کنید (مثلاً از هدر Authorization)
//     req, _ := http.NewRequest("GET", "/", nil)
//     req.Header.Set("Authorization", "Bearer your-jwt-token")
// 
//     // تنظیمات برای تایید JWT
//     validationKeyGetter := func(token *jwt.Token) (interface{}, error) {
//         return secretKey, nil
//     }
// 
//     // خواندن و تایید JWT
//     token, err := request.ParseFromRequest(req, request.AuthorizationHeaderExtractor, validationKeyGetter)
//     if err != nil {
//         fmt.Println("خطا در تایید JWT:", err)
//         return
//     }
// 
//     if token.Valid {
//         // JWT معتبر است
//         claims := token.Claims.(jwt.MapClaims)
//         username := claims["username"].(string)
//         fmt.Println("نام کاربری از JWT:", username)
//     } else {
//         fmt.Println("JWT نامعتبر است.")
//     }
// }
// 

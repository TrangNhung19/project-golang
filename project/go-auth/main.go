package main

import (
	"fmt"
	"net/http"

	authcontroller "github.com/jeypc/go-auth/controllers"
)

func main() {
	http.HandleFunc("/", authcontroller.Index)
	http.HandleFunc("/login", authcontroller.Login)
	http.HandleFunc("/logout", authcontroller.Logout)
	http.HandleFunc("/register", authcontroller.Register)

	fmt.Println("Server jalan di: http://localhost/go-auth")
	http.ListenAndServe(":5000", nil)
}

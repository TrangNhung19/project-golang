package main

import (
	"fmt"
	"net/http"

	accountcontroller "github.com/jeypc/go-auth/controllers"
)

func main() {
	fs := http.FileServer(http.Dir("assets"))
	http.Handle(("/assets/"), http.StripPrefix("/assets", fs))
	fs1 := http.FileServer(http.Dir("vendor"))
	http.Handle(("/vendor/"), http.StripPrefix("/vendor", fs1))

	http.HandleFunc("/", accountcontroller.Index)
	http.HandleFunc("/login", accountcontroller.Login)
	http.HandleFunc("/logout", accountcontroller.Logout)
	http.HandleFunc("/register", accountcontroller.Register)

	fmt.Println(" http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}

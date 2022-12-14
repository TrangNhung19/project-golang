package accountcontroller

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("mysession"))

func Index(response http.ResponseWriter, request *http.Request) {
	tmp, _ := template.ParseFiles("templates/login.html")
	tmp.Execute(response, nil)
}

func Login(response http.ResponseWriter, r *http.Request) {
	request.ParseForm()
	username := request.Form.Get("username")
	password := request.Form.Get("password")
	fmt.Println("username :", username)
	fmt.Println("password : ", password)
	if username == "admin" && password == "123456" {

		sessions, _ := store.Get(request, "mysession")
		session.Values["username"] = username
		http.Redirect(response, request, "/login.html", http.StatusSeeOther)

	} else {
		data := map[string]interface{}{
			"err": "Invalid",
		}
		tmp, _ := template.ParseFiles("templates/login.html")
		tmp.Execute(response, data)
	}
}

func Products(response http.ResponseWriter, request *http.Request) {
	session, _ := store.Get(request, "mysession")
	username := session.Values["username"]
	data := map[string]interface{}{
		"username": username,
	}
	tmp, _ := template.ParseFiles("templates/products.html")
	tmp.Execute(response, nil)
}

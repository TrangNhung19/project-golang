package main

import (
	"html/template"
	"net/http"
	"controllers"
	authcontroller "github.com/jeypc/go-auth/controllers"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))

}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.html", nil)
}

//	func loginHandler(w http.ResponseWriter, r *http.Request) {
//		tmpl.ExecuteTemplate(w, "login.html", nil)
//	}
func productHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "products.html", nil)
}
func cartHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "cart.html", nil)
}
func product_dtHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "product_detail.html", nil)
}
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "about.html", nil)
}
func checkoutHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "checkout.html", nil)
}
func contactHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "contact.html", nil)
}

//	func registerHandler(w http.ResponseWriter, r *http.Request) {
//		tmpl.ExecuteTemplate(w, "register.html", nil)
//	}
func searchHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "search.html", nil)
}

func main() {
	fs := http.FileServer(http.Dir("assets"))
	http.Handle(("/assets/"), http.StripPrefix("/assets", fs))
	fs1 := http.FileServer(http.Dir("vendor"))
	http.Handle(("/vendor/"), http.StripPrefix("/vendor", fs1))

	http.HandleFunc("/", homeHandler)
	// http.HandleFunc("/login.html", loginHandler)
	// http.HandleFunc("/register.html", registerHandler)
	http.HandleFunc("/", authcontroller.Index)
	http.HandleFunc("/login", authcontroller.Login)
	http.HandleFunc("/logout", authcontroller.Logout)
	http.HandleFunc("/register", authcontroller.Register)

	http.HandleFunc("/cart.html", cartHandler)
	http.HandleFunc("/products.html", productHandler)
	http.HandleFunc("/product_dt.html", product_dtHandler)
	http.HandleFunc("/about.html", aboutHandler)
	http.HandleFunc("/contact.html", contactHandler)
	http.HandleFunc("/checkout.html", checkoutHandler)

	http.ListenAndServe(":5000", nil)
}

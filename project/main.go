package main

import (
	"html/template"
	"net/http"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))

}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.html", nil)
}
func loginHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "login.html", nil)
}
func productHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "products.html", nil)
}
func cartHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "cart.html", nil)
}
func main() {
	fs := http.FileServer(http.Dir("assets"))
	http.Handle(("/assets/"), http.StripPrefix("/assets", fs))
	fs1 := http.FileServer(http.Dir("vendor"))
	http.Handle(("/vendor/"), http.StripPrefix("/vendor", fs1))

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/login.html", loginHandler)
	http.HandleFunc("/cart.html", cartHandler)
	http.HandleFunc("/products.html", productHandler)

	http.ListenAndServe(":5000", nil)
}

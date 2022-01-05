package main

import (
	"fmt"
	"net/http"

	"lenslocked/controllers"
	"lenslocked/views"

	"github.com/gorilla/mux"
)

var (
	homeView 			views.View
	contactView		views.View
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(homeView.Render(w, nil))
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(contactView.Render(w, nil))
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "Frequently asked questions")
}

func err404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>We could not find the page you were looking for</p>")
}

func main() {
	homeView = *views.NewView("bootstrap", "views/home.gohtml")
	contactView = *views.NewView("bootstrap", "views/contact.gohtml")

	usersC := controllers.NewUsers()

	r := mux.NewRouter()

	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/signup", usersC.New)
	r.HandleFunc("/faq", faq)
	r.NotFoundHandler = http.HandlerFunc(err404)

	http.ListenAndServe(":3000", r)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
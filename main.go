package main

import (
	"fmt"
	"net/http"

	"lenslocked/controllers"

	"github.com/gorilla/mux"
)

func faq(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Frequently asked questions")
}

func err404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>We could not find the page you were looking for</p>")
}

func main() {
	staticC := *controllers.NewStatic()

	usersC := *controllers.NewUsers()

	r := mux.NewRouter()

	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")

	http.ListenAndServe(":3000", r)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
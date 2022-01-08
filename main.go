package main

import (
	"fmt"
	"net/http"

	"lenslocked/controllers"
	"lenslocked/models"

	"github.com/gorilla/mux"
)

func faq(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Frequently asked questions")
}

func err404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>We could not find the page you were looking for</p>")
}

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "thing"
	dbname = "localdb"
)

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	us := models.NewUserService(psqlInfo)
	defer us.Close()
	//us.DestructiveReset()
	us.AutoMigrate()

	staticC := *controllers.NewStatic()
	usersC := *controllers.NewUsers(us)

	r := mux.NewRouter()

	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.Handle("/signup", usersC.NewView).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	r.Handle("/login", usersC.LoginView).Methods("GET")
	r.HandleFunc("/login", usersC.Login).Methods("POST")

	fmt.Println("Starting server on port 3000...")
	http.ListenAndServe(":3000", r)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
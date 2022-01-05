package controllers

import (
	"fmt"
	"lenslocked/views"
	"net/http"

	"github.com/gorilla/schema"
)

//NewUsers creates a new Users controller.
//will panic if templates not parsed correctly
//only use during initial setup
func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
	}
}

//Users controller
type Users struct {
	NewView		*views.View
}

//New renders the new user page
//
//Get /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

type SignupForm struct {
	Email				string		`schema:"email"`
	Password		string		`schema:"password"`
}

//Create creates a new user account from the signup form on submit
//
//Post /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	dec := schema.NewDecoder()
	var form SignupForm
	if err := dec.Decode(&form, r.PostForm); err != nil {
		panic(err)
	}
	fmt.Fprintln(w, form)
}
package controllers

import (
	"lenslocked/views"
	"net/http"
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
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}
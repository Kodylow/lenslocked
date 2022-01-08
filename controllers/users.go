package controllers

import (
	"fmt"
	"lenslocked/models"
	"lenslocked/views"
	"net/http"
)

//NewUsers creates a new Users controller.
//will panic if templates not parsed correctly
//only use during initial setup
func NewUsers(us *models.UserService) *Users {
	return &Users{
		NewView: 	views.NewView("bootstrap", "users/new"),
		LoginView:	views.NewView("bootstrap", "users/login"),
		us:				us,
	}
}

//Users controller
type Users struct {
	NewView		*views.View
	LoginView	*views.View
	us				*models.UserService
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
	Name				string		`schema:"name"`
	Email				string		`schema:"email"`
	Password		string		`schema:"password"`
}

//Create creates a new user account from the signup form on submit
//
//Post /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var form SignupForm
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	user := models.User{
		Name: form.Name,
		Email: form.Email,
		Password: form.Password,
	}
	if err := u.us.Create(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, user)
}
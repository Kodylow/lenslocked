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

type LoginForm struct {
	Email	string `schema:"email"`
	Password	string	`schema:"password"`
}

//Login verifies email and password then logs user in
//
//Post /login
func (u *Users) Login(w http.ResponseWriter, r *http.Request) {
	form := LoginForm{}
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}

	user, err := u.us.Authenticate(form.Email, form.Password)
	if err != nil {
		switch err {
		case models.ErrNotFound:
			fmt.Fprintln(w, "Invalid email address.")
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
	fmt.Println("User", user.Email)
	cookie := http.Cookie{
		Name: "email",
		Value: user.Email,
	}

	http.SetCookie(w, &cookie)
}

//CookieTest displays cookies set on current user
func (u *Users) CookieTest(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("email")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("In here")
	fmt.Fprintln(w, "Email is: ", cookie.Value)
}
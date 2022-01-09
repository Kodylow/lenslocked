package models

import (
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrNotFound = errors.New("models: resources not found")
	ErrInvalidID = errors.New("models: ID provided was invalid")
	ErrInvalidPW = errors.New("models: incorrect password provided")
)

const userPwdPepper = "satoshi-nakamoto"

// NewUserService creates the singleton of the user service
func NewUserService(connectionInfo string) (*UserService) {
	db, err := gorm.Open("postgres", connectionInfo)
	db.LogMode(true)
	if err != nil {
		return nil
	}
	return &UserService{db: db}
}

// UserService is a gorm db for the user
type UserService struct {
	db *gorm.DB
}

// ByID looks up by id
//1 - user, nil
//2 - nil, ErrNotFound
//3 - nil, otherError
func (us *UserService) ByID(id uint) (*User, error) {
	var user User
	db := us.db.Where("id = ?", id)
	err := first(db, &user)
	return &user, err
}

// ByEmail looks up by email
func (us *UserService) ByEmail(email string) (*User, error) {
	var user User
	db := us.db.Where("email = ?", email)
	err := first(db, &user)
	return &user, err
}

//Autthenticate auths a user with nput email and password
func (us *UserService) Authenticate(email, password string) (*User, error) {
	foundUser, err := us.ByEmail(email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.PasswordHash), []byte(password + userPwdPepper))
	if err != nil {
		switch err {
		case bcrypt.ErrMismatchedHashAndPassword:
			return nil, ErrInvalidPW
		default:
			return nil, err
		}
	}
	return foundUser, nil
}

//first queries using the provided gorm.db
//gets the first item returned and places it in dest
//if nothing found, returns errNotFound
func first(db *gorm.DB, dest interface{}) error {
	err := db.First(dest).Error
	switch err {
	case nil:
		return nil
	case gorm.ErrRecordNotFound:
		return ErrNotFound
	default:
		return err
	}
}

//Create creates a user off the provided data
func (us *UserService) Create(user *User) error {
	pwdBytes := []byte(user.Password + userPwdPepper)
	hashedBytes, err := bcrypt.GenerateFromPassword(pwdBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PasswordHash = string(hashedBytes)
	user.Password = ""
	return us.db.Create(user).Error
}

// Close closes the userservice db connection
func (us *UserService) Close() error {
	us.db.Close()
	return nil
}

//Update updates user with provided user data
func (us *UserService) Update(user *User) error {
	return us.db.Save(user).Error
}

// Delete deletes user with provided id
func (us *UserService) Delete(id uint) error {
	if id == 0 {
		return ErrInvalidID
	}
	user := User{Model: gorm.Model{ID: id}}
	return us.db.Delete(&user).Error
}

// DestructiveReset drops tables and automigrates
func (us *UserService) DestructiveReset() {
	us.db.DropTableIfExists(&User{})
	us.db.AutoMigrate(&User{})
}

func (us *UserService) AutoMigrate() error {
	if err := us.db.AutoMigrate(&User{}).Error; err != nil {
		return err
	}
	return nil
}

type User struct {
	gorm.Model
	Name	string
	Email	string `gorm:"not null;unique_index"`
	Password string `gorm:"-"`
	PasswordHash string `gorm:"not null"`
}


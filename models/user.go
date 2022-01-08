package models

import (
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	ErrNotFound = errors.New("models: resources not found")
)

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
	err := us.db.Where("id = ?", id).First(&user).Error
	switch err {
	case nil:
		return &user, err
	case gorm.ErrRecordNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

//Create creates a user off the provided data
func (us *UserService) Create(user *User) error {
	return us.db.Create(user).Error
}

// Close closes the userservice db connection
func (us *UserService) Close() error {
	us.db.Close()
	return nil
}

// DestructiveReset drops tables and automigrates
func (us *UserService) DestructiveReset() {
	us.db.DropTableIfExists(&User{})
	us.db.AutoMigrate(&User{})
}

type User struct {
	gorm.Model
	Name	string
	Email	string `gorm:"not null;unique_index"`
}
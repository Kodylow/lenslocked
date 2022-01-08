package models

import (
	"fmt"
	"testing"
)

func testingUserService() (*UserService, error) {
	const (
		host = "localhost"
		port = 5432
		user = "postgres"
		password = "thing"
		dbname = "lenslocked_test"
	)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	
	us := NewUserService(psqlInfo)
	us.db.LogMode(false)
	//Clear users table between tests
	us.DestructiveReset()
	return us, nil
}

func TestCreateUser(t *testing.T) {
	us, err := testingUserService()
	if err != nil {
		t.Fatal(err)
	}
	user := User {
		Name:	"Michael Scott",
		Email: "michael@scott.com",
	}
	err = us.Create(&user)
	if err != nil {
		t.Fatal(err)
	}
	if user.ID == 0 {
		t.Errorf("Expected ID > 0. Received %d.", user.ID)
	}
}
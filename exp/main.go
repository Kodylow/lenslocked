package main

import (
	"fmt"
	"lenslocked/models"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

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
	// us.DestructiveReset()
	// user := models.User {
	// 	Name:	"Michael Scott",
	// 	Email: "michael@scott.com",
	// }
	// if err := us.Create(&user); err != nil {
	// 	panic(err)
	// }
	user, err := us.ByID(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
}


// type User struct {
// 		ID int
// 		Name string
// 		Email string
// 	}

// 	var users []User
// 	rows, err := db.Query(`
// 		SELECT id, name, email
// 		FROM users`)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		var user User
// 		err := rows.Scan(&user.ID, &user.Name, &user.Email)
// 		if err != nil {
// 			panic(err)
// 		}
// 		users = append(users, user)
// 	}
// 	fmt.Println(users)

// db, err := sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("Connected!")
// 	defer db.Close()
// 	for i := 1; i <= 6; i ++ {
// 		userID := 1
// 		if i > 3 {
// 			userID = 2
// 		}
// 		amount := i * 100
// 		description := fmt.Sprintf("Test %d", userID)
// 		_, err = db.Query(`
// 			INSERT INTO orders(user_id, amount, description)
// 			VALUES($1, $2, $3)`, userID, amount, description)
// 		if err != nil {
// 			panic(err)
// 		}
// 	// }

// type User struct {
// 	gorm.Model
// 	Name 	string
// 	Email	string `gorm:"not null;unique-index"`
// 	Color	string
// }

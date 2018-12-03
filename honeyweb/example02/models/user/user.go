package user

import (
	"database/sql"
	"fmt"
	"time"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// User defines the user model
type User struct {
	ID      int
	Name    string
	Age     int8
	Created time.Time
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func getDB() *sql.DB {
	db, err := sql.Open("mysql", "root@unix(/tmp/mysql.sock)/test?charset=utf8&parseTime=true")
	checkErr(err)
	return db
}

// GetAll returns all users.
func (u User) GetAll() (users []User) {
	db := getDB()
	defer db.Close()
	fmt.Println("Before query")
	rows, err := db.Query("SELECT id, name, age, created FROM user")
	fmt.Println("After query")
	checkErr(err)

	for rows.Next() {
		fmt.Println("First loop")
		var user User
		err = rows.Scan(&user.ID, &user.Name, &user.Age, &user.Created)
		checkErr(err)
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		checkErr(err)
	}
	rows.Close()

	return
}

// Get gets the user info by ID
func (u User) Get(id int) (user User) {
	return
}

// Save creates a user.
func Save(user User) (u User) {
	return
}

// Delete deletes a user by id.
func Delete(id int) {
	return
}

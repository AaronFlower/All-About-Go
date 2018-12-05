package user

import (
	"database/sql"
	"fmt"
	"log"
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

var db = getDB()

func getDB() *sql.DB {
	db, err := sql.Open("mysql", "root@unix(/tmp/mysql.sock)/test?charset=utf8&parseTime=true")
	checkErr(err)
	return db
}

// GetAll returns all users.
func GetAll() (users []User) {
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
func Get(id int64) (User, error) {
	var u User
	stmt, err := db.Prepare("select id, name, age, created from user where id = ? ")
	defer stmt.Close()
	if err != nil {
		log.Fatal(err)
		return u, err
	}
	err = stmt.QueryRow(id).Scan(&u.ID, &u.Name, &u.Age, &u.Created)
	if err != nil {
		log.Fatal(err)
		return u, err
	}
	return u, nil
}

// Save creates a user.
func Save(user User) (User, error) {
	stmt, err := db.Prepare("INSERT INTO user (id, name, age) VALUES(null, ?, ?)")
	defer stmt.Close()
	if err != nil {
		checkErr(err)
	}
	res, err := stmt.Exec(user.Name, user.Age)
	if err != nil {
		checkErr(err)
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		checkErr(err)
	}
	stmt.Close()
	return Get(lastID)
}

// Delete deletes a user by id.
func Delete(id int) error {
	stmt, err := db.Prepare("DELETE FROM user where id ?")
	defer stmt.Close()
	if err != nil {
		log.Fatal(err)
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

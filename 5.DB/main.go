package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

type user struct {
	uid        int
	username   string
	department string
	created    string
}

func main() {
	db, err := sql.Open("mysql", "root@unix(/tmp/mysql.sock)/test?charset=utf8")
	checkErr(err)
	defer db.Close()

	// insert data
	stmt, err := db.Prepare("INSERT userinfo SET username = ?, department = ?, created = ?")
	checkErr(err)

	res, err := stmt.Exec("Andy", "RD", "2018-11-26")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println("LastInsertId", id)

	// query data
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}
}

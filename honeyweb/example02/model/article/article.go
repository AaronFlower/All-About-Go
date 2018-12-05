package article

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// Article defines the art model
type Article struct {
	ID      int
	Title   string
	Content string
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

// GetAll returns all arts.
func GetAll() (arts []Article) {
	db := getDB()
	defer db.Close()
	rows, err := db.Query("SELECT id, title, content, created FROM article")
	checkErr(err)

	for rows.Next() {
		fmt.Println("First loop")
		var art Article
		err = rows.Scan(&art.ID, &art.Title, &art.Content, &art.Created)
		checkErr(err)
		arts = append(arts, art)
	}
	if err = rows.Err(); err != nil {
		checkErr(err)
	}
	rows.Close()

	return
}

// Get gets the art info by ID
func Get(id int64) (Article, error) {
	var a Article
	stmt, err := db.Prepare("SELECT id, title, content, created FROM article where id = ? ")
	defer stmt.Close()
	if err != nil {
		log.Fatal(err)
		return a, err
	}
	err = stmt.QueryRow(id).Scan(&a.ID, &a.Title, &a.Content, &a.Created)
	if err != nil {
		log.Fatal(err)
		return a, err
	}
	return a, nil
}

// Save creates a art.
func Save(art Article) (Article, error) {
	stmt, err := db.Prepare("INSERT INTO article (id, title, content) VALUES(null, ?, ?)")
	defer stmt.Close()
	if err != nil {
		checkErr(err)
	}
	res, err := stmt.Exec(art.Title, art.Content)
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

// Delete deletes a article by id.
func Delete(id int) error {
	stmt, err := db.Prepare("DELETE FROM article where id ?")
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

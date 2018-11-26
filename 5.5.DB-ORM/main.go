package main

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// User Model Struct
type User struct {
	ID   int
	Name string `orm:"size(100)"`
}

func init() {
	// set default db
	orm.RegisterDataBase("default", "mysql", "root@unix(/tmp/mysql.sock)/test?charset=utf8", 30)

	// register model
	orm.RegisterModel(new(User))

	// create table
	orm.RunSyncdb("default", false, true)
}

func main() {
	o := orm.NewOrm()

	user := User{Name: "Ozil"}

	// insert to table
	id, err := o.Insert(&user)
	fmt.Printf("ID: %d, ERR: %v \n", id, err)

	// update record
	user.Name = "Muset"
	num, err := o.Update(&user)
	fmt.Printf("NUM: %d, ERR: = %+v\n", num, err)

	// read one record
	u := User{ID: user.ID}
	err = o.Read(&u)
	fmt.Printf("User: %+v, ERR: %v\n", u, err)

	// delete the record
	num, err = o.Delete(&u)
	fmt.Printf("Delete operaton Num: %d, ERR: %v\n", num, err)

}

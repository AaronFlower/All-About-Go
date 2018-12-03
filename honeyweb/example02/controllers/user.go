package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/aaronflower/ago/honeyweb/example02/models/user"
	"github.com/aaronflower/honey"
)

// UserController defines a User RESTful resource.
type UserController struct {
	honey.Controller
}

// Get returns the users list
func (c *UserController) Get() {
	v := user.GetAll()
	users, err := json.Marshal(v)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(c.Ct.ResponseWriter, string(users))
}

// Post creates a user.
func (c *UserController) Post() {
	var u user.User
	var age, name string
	r, w := c.Ct.Request, c.Ct.ResponseWriter
	name = r.FormValue("name")
	age = r.FormValue("age")
	if len(name) == 0 || len(age) == 0 {
		fmt.Fprintf(w, "Name and age can't be null!")
		return
	}
	ageInt, err := strconv.Atoi(r.FormValue("age"))
	if err != nil {
		log.Fatal(err)
		fmt.Fprintf(w, "Please input the Age!")
		return
	}
	u.Name = name
	u.Age = int8(ageInt)
	u, err = user.Save(u)
	if err != nil {
		log.Fatal(err)
		fmt.Fprintf(w, "Please input the Age!")
		return
	}
	userByte, err := json.Marshal(u)

	if err != nil {
		log.Fatal(err)
		fmt.Fprintf(w, "Please input the Age!")
		return
	}
	fmt.Fprintf(w, string(userByte))
}

package controllers

import (
	"fmt"

	"github.com/aaronflower/ago/honeyweb/example02/models/user"
	"github.com/aaronflower/honey"
)

// UserController defines a User RESTful resource.
type UserController struct {
	honey.Controller
}

var userModel user.User

// Get returns the users list
func (c *UserController) Get() {
	fmt.Fprintf(c.Ct.ResponseWriter, "Return all users!")
	// user.User.GetAll()
	v := userModel.GetAll()
	fmt.Printf("v = %+v\n", v)
	fmt.Fprintf(c.Ct.ResponseWriter, string(23))
}

// Post creates a user.
func (c *UserController) Post() {
	fmt.Fprintf(c.Ct.ResponseWriter, "Create a user!")
}

package controller

import (
	"fmt"

	"github.com/aaronflower/honey"
)

// UserController defines a User RESTful resource.
type UserController struct {
	honey.Controller
}

// Get returns the users list
func (c *UserController) Get() {
	fmt.Fprintf(c.Ct.ResponseWriter, "Return all users!")
}

// Post creates a user.
func (c *UserController) Post() {
	fmt.Fprintf(c.Ct.ResponseWriter, "Create a user!")
}

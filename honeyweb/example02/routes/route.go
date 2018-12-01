package routes

import (
	"github.com/aaronflower/ago/honeyweb/example02/controllers"
	"github.com/aaronflower/honey"
)

func init() {
	honey.MyApp.Handlers.Add("/user", &controllers.UserController{})
}

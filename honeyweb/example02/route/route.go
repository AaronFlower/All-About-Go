package route

import (
	"github.com/aaronflower/ago/honeyweb/example02/controller"
	"github.com/aaronflower/honey"
)

func init() {
	honey.MyApp.Handlers.Add("/user", &controller.UserController{})
}

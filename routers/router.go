package routers

import (
	"github.com/astaxie/beegae"
	"github.com/jackgris/optima/controllers"
)

func init() {
	beegae.Router("/home", &controllers.MainController{})
	beegae.Router("/user/newuser", &controllers.UserAddController{})
}
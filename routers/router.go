package routers

import (
	"github.com/astaxie/beegae"
	"github.com/jackgris/optima/controllers"
)

func init() {
	beegae.Router("/login", &controllers.LoginController{})
	beegae.Router("/register", &controllers.RegistrationController{})
	beegae.Router("/user/newuser", &controllers.UserAddController{})
	beegae.Router("/", &controllers.MainController{})
}

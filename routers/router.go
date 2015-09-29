package routers

import (
	"github.com/astaxie/beegae"

	"github.com/jackgris/optima/controllers"
)

func init() {
	beegae.Router("/auth/login", &controllers.LoginAuth{})
	beegae.Router("/auth/signup", &controllers.RegisterAuth{})
	beegae.Router("/privatedata", &controllers.PrivateController{})
	beegae.Router("/", &controllers.MainController{})
}

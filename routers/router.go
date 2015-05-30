package routers

import (
	"github.com/astaxie/beegae"
	"github.com/jackgris/optima/controllers"
)

func init() {
	beegae.Router("/", &controllers.MainController{})
}

package controllers

import (
	"github.com/astaxie/beegae"
)

type MainController struct {
	beegae.Controller
}

func (this *MainController) Get() {
	this.Data["Website"] = "optima.com.ar"
	this.Data["Email"] = "optima@gmail.com"
	this.TplNames = "home.tpl"
}

package controllers

import (
	"github.com/astaxie/beegae"
)

type MainController struct {
	beegae.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "index.tpl"
}

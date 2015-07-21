package controllers

import (
	"github.com/astaxie/beegae"
)

type RegistrationController struct {
	beegae.Controller
}

func (this *RegistrationController) Get() {
	this.TplNames = "registration.tpl"
}

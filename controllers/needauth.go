package controllers

import "github.com/astaxie/beegae"

type NeedAuth interface {
	AuthPrivatePlace()
}

type NeedAuthController struct {
	beegae.Controller
}

func (this *NeedAuthController) Prepare() {
	if app, ok := this.AppController.(NeedAuth); ok {
		app.AuthPrivatePlace()
	}
}

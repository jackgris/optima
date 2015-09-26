package controllers

type PrivateController struct {
	NeedAuthController
}

func (this *PrivateController) Get() {
	this.TplNames = "private.tpl"
}

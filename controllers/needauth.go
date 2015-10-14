package controllers

type NeedAuth interface {
	AuthPrivatePlace()
}

type NeedAuthController struct {
	RenderController
}

func (this *NeedAuthController) Prepare() {
	if app, ok := this.AppController.(NeedAuth); ok {
		app.AuthPrivatePlace()
	}
}

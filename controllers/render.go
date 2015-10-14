package controllers

import "github.com/astaxie/beegae"

type RenderController struct {
	beegae.Controller
}

func (this *RenderController) Render() error {
	if _, ok := this.Data["json"].(error); ok {
		this.AppEngineCtx.Errorf("Render error: %v", this.Data["json"])
	}
	this.ServeJson()
	return nil
}

package controllers

import (
	"log"

	"github.com/jackgris/optima/models"
)

type PrivateController struct {
	MiddlewareAuthController
}

func (this *PrivateController) Get() {

	var advertisers []models.Advertiser
	advertisers, err := models.GetAdvertisers(this.AppEngineCtx)
	if err != nil {
		log.Println("Load advertisers error ", err)
	}
	this.Data["json"] = &advertisers
}

func (this *PrivateController) Render() error {
	if _, ok := this.Data["json"].(error); ok {
		this.AppEngineCtx.Errorf("login auth error: %v", this.Data["json"])
	}
	this.ServeJson()
	return nil
}

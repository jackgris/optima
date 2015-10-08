package controllers

import (
	"log"

	"github.com/jackgris/optima/models"
)

type AdvertisersController struct {
	MiddlewareAuthController
}

func (this *AdvertisersController) Get() {

	var advertisers []models.Advertiser
	advertisers, err := models.GetAdvertisers(this.AppEngineCtx)
	if err != nil {
		log.Println("Load advertisers error ", err)
	}
	this.Data["json"] = &advertisers
}

func (this *AdvertisersController) Render() error {
	if _, ok := this.Data["json"].(error); ok {
		this.AppEngineCtx.Errorf("login auth error: %v", this.Data["json"])
	}
	this.ServeJson()
	return nil
}

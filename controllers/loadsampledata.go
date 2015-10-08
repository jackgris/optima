package controllers

import (
	"log"

	"github.com/astaxie/beegae"
	"github.com/jackgris/optima/models"
)

type ErrorLoad struct {
	Error string
}

// The sole purpose of this is to load the information at the beginning
// of the application database to test it
type ExampleDataController struct {
	beegae.Controller
}

func (this *ExampleDataController) Get() {
	err := models.LoadExampleDataAtInit(this.AppEngineCtx)
	if err != nil {
		log.Println("Load data example error ", err)
		e := ErrorLoad{Error: "Error to load data"}
		this.Data["json"] = &e
		return
	}
	this.Data["json"] = &models.Advertiser{}
}

func (this *ExampleDataController) Render() error {
	if _, ok := this.Data["json"].(error); ok {
		this.AppEngineCtx.Errorf("Example data  error: %v", this.Data["json"])
	}
	this.ServeJson()
	return nil
}

package controllers

import (
	"log"

	"github.com/astaxie/beegae"
	"github.com/jackgris/optima/models"
)

type RegisterAuth struct {
	beegae.Controller
}

func (this *RegisterAuth) Post() {

	user, err := decodeUserData(this.Ctx.Input.Request.Body)
	if err != nil {
		this.Data["json"] = err
		log.Fatalln("Error decode user data on register", err)
		return
	}

	log.Println("Registro data ", user)
	token, _ := models.GenerateToken(user.Email)
	user.Token = token
	exist, err := models.CheckExist(user, this.AppEngineCtx)
	if err != nil {
		log.Fatalln("Error at verify user", err)
		this.Data["json"] = err
		return
	}
	if exist {
		this.Data["json"] = "The user already exist"
		log.Fatalln("The user already exist")
		return
	} else {
		models.AddUser(user, this.AppEngineCtx)
	}
	// It's all ok, return the user data on json format
	this.Data["json"] = &user
}

func (this *RegisterAuth) Get() {
	// not implemmented
}

func (this *RegisterAuth) Render() error {
	if _, ok := this.Data["json"].(error); ok {
		this.AppEngineCtx.Errorf("register auth error: %v", this.Data["json"])
	}
	this.ServeJson()
	return nil
}

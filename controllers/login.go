package controllers

import (
	"log"

	"github.com/astaxie/beegae"
	"github.com/jackgris/optima/models"
)

type LoginAuth struct {
	beegae.Controller
}

func (this *LoginAuth) Post() {
	// get the data from the request on json format
	userData, err := DecodeUserData(this.Ctx.Input.Request.Body)
	if err != nil {
		log.Println("LoginAuth: Error decode user data on login", err)
		this.Data["json"] = &models.Token{}
		return
	}

	// get the user data from the datastore
	user, err := models.GetUser(userData.Email, this.AppEngineCtx)
	if err != nil {
		log.Println("LoginAuth: ", err)
		this.Data["json"] = &models.Token{}
		return
	}
	// check if the password is the same
	validPass, err := CheckPassword(user.Pass, user.Salt, userData.Pass)
	if err != nil {
		errorM := "Error checking password"
		log.Println("LoginAuth: ", errorM)
		this.Data["json"] = &models.Token{}
		return
	}
	if !validPass {
		wrong := "Wrong password"
		log.Println("LoginAuth: ", wrong)
		this.Data["json"] = &models.Token{}
		return
	}

	this.Data["json"] = &user.Token
}

func (this *LoginAuth) Get() {
	log.Println("LoginAuth: En el meuserData GET de LoginAuth")
}

func (this *LoginAuth) Render() error {
	if _, ok := this.Data["json"].(error); ok {
		this.AppEngineCtx.Errorf("login auth error: %v", this.Data["json"])
	}
	this.ServeJson()
	return nil
}

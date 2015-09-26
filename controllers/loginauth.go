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
	userData, err := decodeUserData(this.Ctx.Input.Request.Body)
	if err != nil {
		log.Fatalln("Error decode user data on login", err)
		this.Data["json"] = &models.Token{}
		return
	}

	log.Println("AppEngine ", userData)
	// get the user data from the datastore
	user, err := models.GetUser(userData.Email, this.AppEngineCtx)
	if err != nil {
		log.Fatalln("Login", err)
		this.Data["json"] = &models.Token{}
		return
	}
	// check if the password is the same
	if user.Pass != userData.Pass {
		wrong := "Wrong password"
		log.Fatalln("Login", wrong)
		this.Data["json"] = &models.Token{}
		return
	}

	log.Println(len(user.Token.Hash))
	this.Data["json"] = &user
}

func (this *LoginAuth) Get() {
	log.Println("En el meuserData GET de LoginAuth")
}

func (this *LoginAuth) Render() error {
	if _, ok := this.Data["json"].(error); ok {
		this.AppEngineCtx.Errorf("login auth error: %v", this.Data["json"])
	}
	this.ServeJson()
	return nil
}

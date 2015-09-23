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
		this.Data["json"] = err
		log.Fatalln("Error decode user data on login", err)
		return
	}

	log.Println("AppEngine ", userData)
	// get the user data from the datastore
	user, err := models.GetUser(userData.Email, this.AppEngineCtx)
	if err != nil {
		this.Data["json"] = err
		log.Fatalln("Login", err)
		return
	}
	// check if the password is the same
	if user.Pass != userData.Pass {
		wrong := "Wrong password"
		this.Data["json"] = wrong
		log.Fatalln("Login", wrong)
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

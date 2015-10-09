package controllers

import (
	"log"

	"encoding/hex"
	"github.com/astaxie/beegae"
	"github.com/golang/crypto/bcrypt"
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
	pass, err := hex.DecodeString(user.Pass)
	if err != nil {
		wrong := "Error decodeString"
		log.Println("LoginAuth: ", wrong)
		this.Data["json"] = &models.Token{}
		return
	}
	log.Println("comparando")
	err = bcrypt.CompareHashAndPassword(pass, []byte(userData.Pass))
	if err != nil {
		wrong := "Wrong password"
		log.Println("LoginAuth: ", wrong)
		this.Data["json"] = &models.Token{}
		return
	}
	log.Println("termino")

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

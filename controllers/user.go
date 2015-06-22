package controllers

import (
	//"github.com/Jackgris/optima/models"
	//"github.com/astaxie/beegae
	"log"
)
type UserAddController MainController
type UserVerifyController MainController

// this call the form to add a new user
func (this *UserAddController) Get(){
	this.TplNames= "user/newuser.tpl"
}

// in this we get the values and reate a new User and persist it
func (this *UserVerifyController) Post() {
	//c :=  &this.AppEngineCtx
	username := this.GetString("username")
	password := this.GetString("password")
	
	log.Println("user ", username)
	log.Println("password ", password)
}
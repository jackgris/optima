package controllers

import (
	"log"
	"strings"

	"github.com/astaxie/beego/validation"
	"github.com/jackgris/optima/models"
)

type LoginController MainController

// Get() show login form
func (this *LoginController) Get() {
	this.TplNames = "login.tpl"
}

// Post() get user input, check the data  and shows us the answer
func (this *LoginController) Post() {

	email := this.GetString("user_email")
	pass := this.GetString("user_password")
	typeuser := this.GetString("type_user")

	user := userLogin{
		Email: email,
		Pass:  pass,
		Type:  typeuser,
	}

	valid := validation.Validation{}
	v, err := valid.Valid(&user)
	if err != nil {
		// handle error
		log.Println("Error valid user", err)
	}
	if !v {
		log.Println("Validation not pass")
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
		this.TplNames = "login.tpl"
		return
	}

	login := &models.User{
		Pass:  user.Pass,
		Email: user.Email,
	}

	exist, err := models.CheckExist(login, this.AppEngineCtx)
	if err != nil {
		log.Fatalln("Error to check if user exist")
		this.TplNames = "login.tpl"
		return
	}

	if exist {
		// FIX need create a new view if the user already exist
		this.TplNames = "login.tpl"
	}

	this.TplNames = "login.tpl"
}

// user useful struct to validate user input
type userLogin struct {
	Email string `valid:"Email; MaxSize(100)"` // Need to be a valid Email address and no more than 100 characters.
	Pass  string `valid:"AlphaNumeric; MaxSize(10)"`
	Type  string `valid:"Required"`
}

// Valid function that will help us to compare the data with some restrictions that we have created
func (u *userLogin) Valid(v *validation.Validation) {

	if strings.Index(u.Email, "admin") != -1 {
		v.SetError("Name", "Can't contain 'admin' in Name")
	}
}

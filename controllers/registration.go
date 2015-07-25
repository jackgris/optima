package controllers

import (
	"log"
	"strings"

	"github.com/astaxie/beegae"
	"github.com/astaxie/beego/validation"
	"github.com/jackgris/optima/models"
)

type RegistrationController struct {
	beegae.Controller
}

// Get() show registration form
func (this *RegistrationController) Get() {
	this.TplNames = "registration.tpl"
}

// Post() get user input, check the data, save user data and shows us the answer
func (this *RegistrationController) Post() {

	name := this.GetString("user_name")
	email := this.GetString("user_email")
	pass := this.GetString("user_password")
	passagain := this.GetString("user_password_again")
	age, err := this.GetInt("user_age")
	if err != nil {
		log.Println("Error getting age", err)
		this.TplNames = "registration.tpl"
		return
	}
	biography := this.GetString("user_bio")
	job := this.GetString("user_job")
	interests := this.GetStrings("user_interest")

	user := &user{
		Name:      name,
		Email:     email,
		Pass:      pass,
		PassAgain: passagain,
		Age:       age,
		Bio:       biography,
		Job:       job,
		Interest:  interests,
	}

	valid := validation.Validation{}
	v, err := valid.Valid(user)
	if err != nil {
		// handle error
		log.Println("Error valid user", err)
	}
	if !v {
		log.Println("Validation not pass")
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
		this.TplNames = "registration.tpl"
		return
	}

	register := &models.User{
		Email: email,
	}

	exist, err := models.CheckExist(register, this.AppEngineCtx)
	if err != nil {
		log.Fatalln("Error: check if user exist, ", err.Error())
		this.TplNames = "registration.tpl"
		return
	}

	if exist {
		log.Println("The user already exist")
		// FIX need create a new view if the user already exist
		this.TplNames = "registration.tpl"
		return
	}

	_, err = models.AddUser(register, this.AppEngineCtx)
	if err != nil {
		log.Fatalln("Error created new user")
		// FIX need create a new view if the user already exist
		this.TplNames = "registration.tpl"
		return
	}

	this.TplNames = "registration.tpl"
}

// user useful struct to validate user input
type user struct {
	Name      string   `valid:"Required"`
	Email     string   `valid:"Email; MaxSize(100)"` // Need to be a valid Email address and no more than 100 characters.
	Pass      string   `valid:"AlphaNumeric; MaxSize(10)"`
	PassAgain string   `valid:"AlphaNumeric; MaxSize(10)"`
	Age       int      `valid:"Range(18, 100)"` // 1 <= Age <= 140, only valid in this range
	Bio       string   `valid:"MaxSize(200)"`
	Job       string   `valid:"Required"`
	Interest  []string `valid:"MinSize(1)"`
}

// Valid function that will help us to compare the data with some restrictions that we have created
func (u *user) Valid(v *validation.Validation) {

	if strings.Index(u.Name, "admin") != -1 {
		v.SetError("Name", "Can't contain 'admin' in Name")
	} else if !strings.EqualFold(u.Pass, u.PassAgain) {
		v.SetError("Password", "The entered passwords are not the same")
	}
}

package controllers

import (
	"fmt"
	"log"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/jackgris/optima/models"
)

type PrivateAuthController struct {
	NeedAuthController
}

func (this *PrivateAuthController) AuthPrivatePlace() {

	s := strings.SplitN(this.Ctx.Request.Header.Get("Authorization"), " ", 2)
	if len(s) != 2 || s[0] != "Bearer" {
		log.Println("PrivateAuth: Hasn't authorization header")
		this.Redirect("/", 302)
		return
	} else {
		token := strings.SplitN(this.Ctx.Request.Header.Get("Authorization"), " ", 2)
		payload, err := jwt.Parse(token[1], func(token *jwt.Token) (interface{}, error) {
			// Validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(models.Privatekey), nil
		})

		switch err.(type) {
		case nil:
			if !payload.Valid {
				log.Println("PrivateAuth: Invalid payload", err)
				this.Redirect("/", 302)
				return
			}
		case *jwt.ValidationError:
			vErr := err.(*jwt.ValidationError)
			switch vErr.Errors {
			case jwt.ValidationErrorExpired:
				log.Println("PrivateAuth: Token expired", err)
				this.Redirect("/", 302)
				return
			default:
				log.Println("PrivateAuth: Error validation", err)
				this.Redirect("/", 302)
				return
			}
		default:
			if err != nil {
				log.Println("PrivateAuth: Error payload", err)
				this.Redirect("/", 302)
				return
			}
		}

		if email, ok := payload.Claims["sub"].(string); !ok {
			log.Println("PrivateAuth: Error get email user from token", err)
			this.Redirect("/", 302)
		} else {
			ok, err := models.CheckExist(&models.User{Email: email}, this.AppEngineCtx)
			if err != nil {
				log.Println("PrivateAuth: Error check user", err)
				this.Redirect("/", 302)
			}
			if !ok {
				log.Println("PrivateAuth: The user not exist")
				this.Redirect("/", 302)
			}
		}
	}
}

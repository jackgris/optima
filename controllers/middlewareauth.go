package controllers

import (
	"fmt"
	"log"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/jackgris/optima/models"
)

type MiddlewareAuthController struct {
	NeedAuthController
}

func (this *MiddlewareAuthController) AuthPrivatePlace() {
	// Check if has authentication header, if not, redirect to the main page
	s := strings.SplitN(this.Ctx.Request.Header.Get("Authorization"), " ", 2)
	if len(s) != 2 || s[0] != "Bearer" {
		log.Println("MiddlewareAuth: Hasn't authorization header")
		this.Redirect("/", 302)
		return
	} else {
		// Check if has an valid token
		token := strings.SplitN(this.Ctx.Request.Header.Get("Authorization"), " ", 2)
		payload, err := jwt.Parse(token[1], func(token *jwt.Token) (interface{}, error) {
			// Validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(models.Privatekey), nil
		})
		// Verify diferents errors type, than can made from the token
		switch err.(type) {
		case nil:
			if !payload.Valid {
				log.Println("MiddlewareAuth: Invalid payload", err)
				this.Redirect("/", 302)
				return
			}
		case *jwt.ValidationError:
			vErr := err.(*jwt.ValidationError)
			switch vErr.Errors {
			case jwt.ValidationErrorExpired:
				log.Println("MiddlewareAuth: Token expired", err)
				this.Redirect("/", 302)
				return
			default:
				log.Println("MiddlewareAuth: Error validation", err)
				this.Redirect("/", 302)
				return
			}
		default:
			if err != nil {
				log.Println("MiddlewareAuth: Error payload", err)
				this.Redirect("/", 302)
				return
			}
		}

		// Get the email from the token
		if email, ok := payload.Claims["sub"].(string); !ok {
			log.Println("MiddlewareAuth: Error get email user from token", err)
			this.Redirect("/", 302)
		} else {
			// Veify if exist an user with that email
			ok, err := models.CheckExist(&models.User{Email: email}, this.AppEngineCtx)
			if err != nil {
				log.Println("MiddlewareAuth: Error check user", err)
				this.Redirect("/", 302)
			}
			if !ok {
				log.Println("MiddlewareAuth: The user not exist")
				this.Redirect("/", 302)
			}
		}
	}
}

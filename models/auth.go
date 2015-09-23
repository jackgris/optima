package models

import (
	"errors"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	Privatekey    = "1234"
	tokenduration = 24
)

func GenerateToken(user string) (Token, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims["exp"] = time.Now().Add(time.Hour * tokenduration)
	token.Claims["iat"] = time.Now().Unix()
	token.Claims["sub"] = user
	hash, err := token.SignedString([]byte(Privatekey))
	// var t Token
	t := Token{}

	if err != nil {
		log.Fatalln(errors.New("Error generating hash token" + err.Error()))
		return t, errors.New("Error generating hash token" + err.Error())
	}
	t.Hash = hash
	t.Expire = tokenduration
	return t, nil
}

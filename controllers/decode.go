package controllers

import (
	"encoding/json"
	"github.com/jackgris/optima/models"

	"golang.org/x/crypto/bcrypt"
	"io"
	"log"
	"math/rand"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const lenSalt = 10

func DecodeUserData(r io.ReadCloser) (*models.User, error) {
	defer r.Close()
	var u models.User
	err := json.NewDecoder(r).Decode(&u)
	return &u, err
}

/*
	checking the input password be the user's password
*/
func CheckPassword(hash []byte, salt []byte, pass []byte) (valid bool, err error) {
	log.Println("comparando")
	// use another algorithm that append method
	passSalt := append(pass[:], salt[:]...)
	err = bcrypt.CompareHashAndPassword(hash, passSalt)
	log.Println("termino")
	if err != nil {
		wrong := "the password isn't the same"
		log.Println("CheckPass: ", wrong)
		return false, err
	}

	return true, nil

}

func GenerateRandomSalt() []byte {
	rand.Seed(time.Now().UTC().UnixNano())
	b := make([]byte, lenSalt)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return b
}

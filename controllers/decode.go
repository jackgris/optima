package controllers

import (
	"encoding/json"

	"github.com/jackgris/optima/models"

	"io"
	"log"
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const lenSalt = 10

// This decode input data on json format and return the User with that data
func DecodeUserData(r io.ReadCloser) (*models.User, error) {
	defer r.Close()
	var u models.User
	err := json.NewDecoder(r).Decode(&u)
	return &u, err
}

// This decode input data on json format and return the Advertiser with that data
func DecodeAdvertiserData(r io.ReadCloser) (*models.Advertiser, error) {
	defer r.Close()
	var u models.Advertiser
	err := json.NewDecoder(r).Decode(&u)
	return &u, err
}

// Checking the input password be the user's password
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

// Create a salt needed for create the hash used for register and login user
func GenerateRandomSalt() []byte {
	rand.Seed(time.Now().UTC().UnixNano())
	b := make([]byte, lenSalt)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return b
}

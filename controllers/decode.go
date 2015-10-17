package controllers

import (
	"encoding/json"
	"io"

	"github.com/jackgris/optima/models"
)

func DecodeUserData(r io.ReadCloser) (*models.User, error) {
	defer r.Close()
	var u models.User
	err := json.NewDecoder(r).Decode(&u)
	return &u, err
}

func DecodeAdvertiserData(r io.ReadCloser) (*models.Advertiser, error) {
	defer r.Close()
	var u models.Advertiser
	err := json.NewDecoder(r).Decode(&u)
	return &u, err
}

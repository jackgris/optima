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

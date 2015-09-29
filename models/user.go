package models

import (
	"errors"

	"appengine"
	"appengine/datastore"
)

// Fetch the user data from the datastore
func GetUser(email string, c appengine.Context) (User, error) {
	q := datastore.NewQuery("User").Filter("Email =", email)
	var users []User
	if _, err := q.GetAll(c, &users); err != nil {
		err = errors.New("on GetUser All" + err.Error())
		return User{}, err
	}

	if len(users) < 1 {
		err := errors.New("on GetUser the user doesn't exist")
		return User{}, err
	}

	user := users[0]
	user.Token, _ = GenerateToken(email)
	return user, nil
}

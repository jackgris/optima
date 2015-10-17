package models

import (
	"errors"
	"log"
	"strings"
	"time"

	"appengine"
	"appengine/datastore"
)

type User struct {
	Id      int64
	Name    string `json:"name"`
	Pass    []byte `json:"password"`
	Email   string `json:"email"`
	Salt    []byte `json: salt`
	Token   Token  `json:"token"`
	Created time.Time
}

// Return the key for use the datastore
func DefaultUserKey(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, "User", "default", 0, nil)
}

// AddUser add a user to datastore
func AddUser(u *User, c appengine.Context) (*datastore.Key, error) {

	u.Created = time.Now()
	key := datastore.NewIncompleteKey(c, "User", DefaultUserKey(c))
	_, err := datastore.Put(c, key, u)
	if err != nil {
		log.Println("AddUser error", err)
	}
	return key, err
}

// CheckExist verifies that the user has not been created earlier
func CheckExist(u *User, c appengine.Context) (bool, error) {

	q := datastore.NewQuery("User").Filter("Email =", u.Email)
	var users []User
	_, err := q.GetAll(c, &users)
	if err != nil {
		return false, errors.New("Error veryfing if user exist," + err.Error())
	}
	if len(users) < 1 {
		return false, nil
	} else {
		for _, us := range users {
			if strings.EqualFold(us.Email, u.Email) {
				return true, nil
			}
		}
	}
	return false, nil
}

// Recreate the token
func (u *User) RefreshToken() {
	token, _ := GenerateToken(u.Email)
	u.Token = token
}

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

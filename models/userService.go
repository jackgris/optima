package models

import (
	"errors"

	"appengine"
	"appengine/datastore"

	"log"
	"reflect"
)

type User struct {
	Name  string
	Pass  string
	Email string
}

// AddUser add a user to datastore
func AddUser(u User, c *appengine.Context) (*datastore.Key, error) {
	//c := appengine.NewContext(rq)
	key := datastore.NewIncompleteKey(*c, "User", nil)
	_, err := datastore.Put(*c, key, &u)
	log.Println("AddUser receibed an object of type", reflect.TypeOf(u))
	return key, err
}

// CheckExist verifies that the user has not been created earlier
func CheckExist(u *User, c appengine.Context) (bool, error) {

	q := datastore.NewQuery("User").
		Filter("Email =", u.Email)
	var users []User
	_, err := q.GetAll(c, &users)
	if err != nil {
		return false, errors.New("Error query datastore when check exists user: " + err.Error())
	}

	if len(users) < 1 {
		return false, nil
	}

	return true, nil
}

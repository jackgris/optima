package models

import (
	"appengine"
	"appengine/datastore"

	"log"
	"reflect"
)

type User struct {
	name_user string
	password  string
}

// add a user to datastore
func AddUser(u User, c *appengine.Context) (*datastore.Key, error) {
	//c := appengine.NewContext(rq)
	key := datastore.NewIncompleteKey(*c, "User", nil)
	_, err := datastore.Put(*c, key, &u)
	log.Println("AddUser receibed an object of type", reflect.TypeOf(u))
	return key, err
}

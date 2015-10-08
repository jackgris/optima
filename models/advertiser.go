package models

import (
	"encoding/json"
	"errors"
	"log"

	"appengine"
	"appengine/datastore"
)

type Advertiser struct {
	Age       int      `json:"age"`
	Id        string   `json:"id"`
	Name      string   `json:"name"`
	Sex       string   `json:"sex"`
	Nse       string   `json:"nse"`
	Coverage  string   `json:"coverage"`
	Interets  []string `json:"interests"`
	Category  string   `json:"category"`
	Budget    int      `json:"budget"`
	Objetives string   `json:"objetives"`
}

func DefaultAdvertiserKey(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, "Advertiser", "default", 0, nil)
}

// Fetch all the advertisers data from the datastore
func GetAdvertisers(c appengine.Context) ([]Advertiser, error) {

	q := datastore.NewQuery("Advertiser")
	var advertisers []Advertiser
	if _, err := q.GetAll(c, &advertisers); err != nil {
		err = errors.New("on Get Advertisers" + err.Error())
		return make([]Advertiser, 0, 0), err
	}

	return advertisers, nil
}

// Load and save the data from the advertiser on the datastore
func LoadAdvertiser(a *Advertiser, c appengine.Context) (*datastore.Key, error) {

	key := datastore.NewIncompleteKey(c, "Advertiser", DefaultAdvertiserKey(c))
	_, err := datastore.Put(c, key, a)
	if err != nil {
		log.Println("LoadAdvertiser error", err)
	}

	return key, err
}

// This function only was need for load example data, for test the aplicacion
func LoadExampleDataAtInit(c appengine.Context) error {
	a := []byte(advertisersData)
	var advertisers []Advertiser
	err := json.Unmarshal(a, &advertisers)
	if err != nil {
		log.Println("Load Example: Error to unmarshal json advertisers", err)
		return err
	}

	for _, ad := range advertisers {
		_, _ = LoadAdvertiser(&ad, c)
	}

	return nil
}

// Example data
const (
	advertisersData = `[
    {
        "age": 1, 
        "id": "motorola-xoom-with-wi-fi", 
        "name": "Motorola XOOOOOOOOM\u2122 with Wi-Fi", 
        "sex": "f",
        "nse": "A",
        "coverage": "test",
        "interests": [""],
        "category": "C",
        "budget": 1000,
        "objetives": "testing"
    },
    {
        "age": 2, 
        "id": "motorola-xoom-with-wi-fi", 
        "name": "Motorola XOOM\u2122 with Wi-Fi", 
        "sex": "f",
        "nse": "A",
        "coverage": "test",
        "interests": [""],
        "category": "C",
        "budget": 1000,
        "objetives": "testing"
    }, 
    {
        "age": 5, 
        "id": "motorola-xoom-with-wi-fi", 
        "name": "AMotorola XOOM\u2122 with Wi-Fi", 
        "sex": "f",
        "nse": "A",
        "coverage": "test",
        "interests": [""],
        "category": "C",
        "budget": 1000,
        "objetives": "testing"
    }, 
    {
        "age": 8, 
        "id": "motorola-xoom-with-wi-fi", 
        "name": "BMotorola XOOM\u2122 with Wi-Fi", 
        "sex": "f",
        "nse": "A",
        "coverage": "test",
        "interests": [""],
        "category": "C",
        "budget": 1000,
        "objetives": "testing"
    }, 
    {
        "age": 0, 
        "id": "motorola-xoom-with-wi-fi", 
        "name": "CMotorola XOOM\u2122 with Wi-Fi", 
        "sex": "f",
        "nse": "A",
        "coverage": "test",
        "interests": [""],
        "category": "C",
        "budget": 1000,
        "objetives": "testing"
    }, 
    {
        "age": 9, 
        "id": "motorola-xoom-with-wi-fi", 
        "name": "DMotorola XOOM\u2122 with Wi-Fi", 
        "sex": "f",
        "nse": "A",
        "coverage": "test",
        "interests": [""],
        "category": "C",
        "budget": 1000,
        "objetives": "testing"
    }, 
    {
        "age": 10, 
        "id": "motorola-xoom-with-wi-fi", 
        "name": "EMotorola XOOM\u2122 with Wi-Fi", 
        "sex": "f",
        "nse": "A",
        "coverage": "test",
        "interests": [""],
        "category": "C",
        "budget": 1000,
        "objetives": "testing"
    }, 
    {
        "age": 50, 
        "id": "motorola-xoom-with-wi-fi", 
        "name": "FMotorola XOOM\u2122 with Wi-Fi", 
        "sex": "f",
        "nse": "A",
        "coverage": "test",
        "interests": [""],
        "category": "C",
        "budget": 1000,
        "objetives": "testing"
    }, 
    {
        "age": 0, 
        "id": "motorola-xoom-with-wi-fi", 
        "name": "Motorola XOOM\u2122 with Wi-Fi", 
        "sex": "f",
        "nse": "A",
        "coverage": "test",
        "interests": [""],
        "category": "C",
        "budget": 1000,
        "objetives": "testing"
    },  
    {
        "age": 0, 
        "id": "motorola-xoom-with-wi-fi", 
        "name": "Motorola XOOM\u2122 with Wi-Fi", 
        "sex": "f",
        "nse": "A",
        "coverage": "test",
        "interests": [""],
        "category": "C",
        "budget": 1000,
        "objetives": "testing"
    }
]`
)

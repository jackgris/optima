package controllers

import (
	"encoding/json"
	"log"

	"github.com/jackgris/optima/models"
)

type PrivateController struct {
	MiddlewareAuthController
}

func (this *PrivateController) Get() {

	a := []byte(advertisersData)
	var advertisers []models.Advertiser
	err := json.Unmarshal(a, &advertisers)
	if err != nil {
		log.Println("Private: Error to unmarshal json advertisers", err)
		this.Data["json"] = &models.Advertiser{}
	}
	this.Data["json"] = &advertisers
}

func (this *PrivateController) Render() error {
	if _, ok := this.Data["json"].(error); ok {
		this.AppEngineCtx.Errorf("login auth error: %v", this.Data["json"])
	}
	this.ServeJson()
	return nil
}

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

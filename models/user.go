package models

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"appengine"
	"appengine/datastore"

	jwt "github.com/dgrijalva/jwt-go"
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

///////////////////////////////////////////////////////////////////
type UserAuth struct {
	Name     string
	Pass     string
	Id       string
	Response http.ResponseWriter
}

func (u *UserAuth) CheckAuth(r *http.Request, w http.ResponseWriter) bool {
	s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(s) != 2 || s[0] != "Basic" {
		u.RequireHeaderAuth()
		return false
	}

	token := strings.SplitN(r.Header.Get("Authorization"), " ", 1)
	payload, err := jwt.Parse(token[0], func(token *jwt.Token) (interface{}, error) {
		// Validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return Privatekey, nil //"1234", nil // debo usar la contante Privatekey de auth
	})

	if err != nil {
		u.RequireHeaderAuth()
		return false
	}

	if !payload.Valid {
		u.RequireAuthExpire()
		return false
	}

	u.Id = payload.Claims["sub"].(string)
	return true
}

func (u *UserAuth) RequireAuthExpire() {
	u.Response.Header().Set("WWW-Authenticate", `Basic realm="`+u.Name+`"`)
	u.Response.WriteHeader(401)
	u.Response.Write([]byte("401 The token has expired\n"))
}

func (u *UserAuth) RequireHeaderAuth() {
	u.Response.Header().Set("WWW-Authenticate", `Basic realm="`+u.Name+`"`)
	u.Response.WriteHeader(403)
	u.Response.Write([]byte("403 No authorization header\n"))
}

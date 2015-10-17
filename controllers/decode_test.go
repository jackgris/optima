package controllers_test

import (
	"bytes"
	"io"
	"net/http"

	"github.com/jackgris/optima/controllers"
	"github.com/jackgris/optima/models"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("Decode", func() {

	Expect := gomega.Expect
	Describe := ginkgo.Describe
	It := ginkgo.It

	Describe("Decode input from request", func() {
		It("Should get the same user info from the json on a request", func() {
			u := models.User{}
			u.Name = "fake"
			u.Email = "fake@gmail.com"
			data := `{"name":"fake","email":"fake@gmail.com"}`
			r := http.Request{}
			r.Body = myCloser{bytes.NewBufferString(data)}
			u2, err := controllers.DecodeUserData(r.Body)
			Expect(err).To(gomega.BeNil())
			Expect(u.Name).To(gomega.Equal(u2.Name))
			Expect(u.Email).To(gomega.Equal(u2.Email))
		})

		It("Should get advertiser data from a json request", func() {
			a := models.Advertiser{}
			a.Name = "fake"
			a.Sex = "s"
			data := `{"name":"fake", "sex":"m"}`
			r := http.Request{}
			r.Body = myCloser{bytes.NewBufferString(data)}
			a2, err := controllers.DecodeAdvertiserData(r.Body)
			Expect(err).To(gomega.BeNil())
			Expect(a.Name).To(gomega.Equal(a2.Name))
			Expect(a.Sex).To(gomega.Equal(a2.Sex)) // Need finish write test
		})
	})
})

var _ = ginkgo.Describe("checking password user", func() {
	Expect := gomega.Expect
	Describe := ginkgo.Describe
	It := ginkgo.It

	Describe("Comparing the password input", func() {
		It("Should get the same input password with database hash password ", func() {
			data := `{"password":"1234","email":"ema@mail.com"}`
			r := http.Request{}
			r.Body = myCloser{bytes.NewBufferString(data)}
			u, err := controllers.DecodeUserData(r.Body)
			Expect(err).To(gomega.BeNil())

			salt := []byte{107, 86, 82, 84, 122, 76, 120, 87, 104, 70}
			hash := []byte{36, 50, 97, 36, 49, 48, 36, 54, 85, 88, 47, 118, 103, 69, 66, 106, 97, 67, 106, 83, 86, 105, 54, 85, 117, 117, 88, 52, 101, 50, 105, 121, 77, 67, 71, 111, 47, 79, 67, 52, 75, 86, 79, 116, 67, 73, 98, 97, 76, 119, 119, 54, 66, 116, 106, 116, 99, 80, 47, 79}
			validPass, err := controllers.CheckPassword(hash, salt, u.Pass)
			Expect(err).To(gomega.BeNil())
			Expect(validPass).To(gomega.BeTrue())

		})

	})

})

type myCloser struct {
	io.Reader
}

func (myCloser) Close() error { return nil }

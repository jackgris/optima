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

type myCloser struct {
	io.Reader
}

func (myCloser) Close() error { return nil }

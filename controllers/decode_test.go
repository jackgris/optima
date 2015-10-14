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
		It("Should get the same info from the json on a request", func() {
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
	})
})

var _ = ginkgo.Describe("checking password user", func() {
	Expect := gomega.Expect
	Describe := ginkgo.Describe
	It := ginkgo.It

	Describe("Comparing the password input", func() {
		It("Should get the same input password with database hash password ", func() {

		})

	})

})

type myCloser struct {
	io.Reader
}

func (myCloser) Close() error { return nil }

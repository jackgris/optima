package models_test

import (
	"strings"
	"time"

	"github.com/jackgris/optima/models"

	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("Auth", func() {

	Expect := gomega.Expect
	Describe := ginkgo.Describe
	It := ginkgo.It

	Describe("Generating token", func() {

		email := "fake@fake.com"
		token, err := models.GenerateToken(email)

		It("Should generate a token", func() {
			Expect(err).To(gomega.BeNil())
			n := strings.SplitN(token.Hash, ".", 3)
			Expect(len(n)).To(gomega.Equal(3))
		})

		It("Generate token from string hash", func() {
			t, err := models.ParseToken(token.Hash)
			Expect(err).To(gomega.BeNil())
			e := t.Claims["sub"].(string)
			Expect(e).To(gomega.Equal(email))
			tNow := t.Claims["iat"]
			timeToken := float64(time.Now().Unix())
			Expect(tNow).To(gomega.Equal(timeToken))
		})
	})
})

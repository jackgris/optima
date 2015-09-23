package tests_test

import (
	"github.com/jackgris/optima/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Login", func() {

	It("Should generate a token", func() {
		email := "fake@fake.com"
		token, err := models.GenerateToken(email)
		Expect(err).To(BeNil())
		Expect(len(token.Hash)).To(Equal(193))
	})
})

package helpers_test

import (
	"google-scraper/helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Password", func() {
	Describe("HashPassword", func() {
		Context("given a string", func() {
			It("returns the hashed version of the string", func() {
				Expect(helpers.HashPassword("Secret")).To(ContainSubstring("$"))
			})
		})
	})
})

package helpers_test

import (
	"google-scraper/helpers"
	. "google-scraper/tests"

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

	Describe("#VerifyPasswordHash", func() {
		Context("given the password and digest are valid", func() {
			It("returns no error", func() {
				password := "secret"
				user := FabricateUser("John", "john@example.com", password)
				err := helpers.VerifyPasswordHash(password, user.HashedPassword)

				Expect(err).To(BeNil())
			})
		})

		Context("given the password and digest are INVALID", func() {
			It("returns an error", func() {
				password := "secret"
				FabricateUser("John", "john@example.com", password)
				err := helpers.VerifyPasswordHash(password, "wrong-digest")

				Expect(err).NotTo(BeNil())
			})
		})
	})

	AfterEach(func() {
		TruncateTables("users")
	})
})

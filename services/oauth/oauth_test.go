package oauth_test

import (
	_ "google-scraper/initializers"
	"google-scraper/services/oauth"
	. "google-scraper/tests"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("OauthService", func() {
	Describe("#PasswordAuthorizationHandler", func() {
		Context("given the valid credential", func() {
			It("returns user ID", func() {
				email := "john@example.com"
				password := "secret"
				FabricateUser("John", email, password)

				Id, err := oauth.PasswordAuthorizationHandler(email, password)
				if err != nil {
					Fail("Authentication failed")
				}

				Expect(Id).NotTo(BeNil())
			})
		})

		Context("given the INVALID credential", func() {
			It("returns an error", func() {
				FabricateUser("John", "john@example.com", "secret")

				_, err := oauth.PasswordAuthorizationHandler("john@example.com", "invalid")

				Expect(err.Error()).To(Equal("invalid_client"))
			})
		})
	})
	AfterEach(func() {
		TruncateTables("users")
	})
})

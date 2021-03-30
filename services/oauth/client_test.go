package oauth_test

import (
	_ "google-scraper/initializers"
	"google-scraper/services/oauth"
	. "google-scraper/tests"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("OauthService", func() {
	Describe("#GenerateOauthClient", func() {
		It("returns oauth client", func() {
			user := FabricateUser("John", "john@example.com", "secret")

			client, err := oauth.GenerateOauthClient(user.Id)
			if err != nil {
				Fail("Creating oauth client failed")
			}

			Expect(client.ID).NotTo(BeNil())
			Expect(client.Secret).NotTo(BeNil())
			Expect(client.UserID).NotTo(BeNil())
		})
	})
})

package serializers_test

import (
	"google-scraper/serializers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TokenSerializer", func() {
	Describe("#GetTokenResponse", func() {
		It("returns token response object", func() {
			json := `{
						"access_token": "access-token",
						"expires_in": 7200,
						"refresh_token": "refresh-token",
						"token_type": "Bearer",
					}`

			response := serializers.GetTokenResponse(json)

			Expect(response.AccessToken).To(Equal("access-token"))
			Expect(response.ExpiresIn).To(Equal(int64(7200)))
			Expect(response.RefreshToken).To(Equal("refresh-token"))
			Expect(response.TokenType).To(Equal("Bearer"))
		})
	})
})

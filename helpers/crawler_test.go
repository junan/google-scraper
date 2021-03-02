package helpers_test

import (
	. "google-scraper/helpers"
	. "google-scraper/constants"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Crawler", func() {
	Describe("#BuildSearchUrl", func() {
		Context("given the root URL is valid", func() {
			It("returns the correct URL", func() {
				url, err := BuildSearchUrl("Ruby", GoogleSearchBaseUrl)
				if err != nil {
					Fail("Building search URL failed: " + err.Error())
				}

				Expect(url).To(Equal("https://www.google.com/search?hl=en&lr=lang_en&q=Ruby"))
			})
		})

		Context("given the root URL is invalid", func() {
			It("returns an error", func() {
				url, err := BuildSearchUrl("Ruby", "http://google:com")

				Expect(err.Error()).To(Equal("parse \"http://google:com\": invalid port \":com\" after host"))
				Expect(url).To(Equal(""))
			})
		})
	})
})

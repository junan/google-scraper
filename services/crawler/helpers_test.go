package crawler_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "google-scraper/services/crawler"
)

var _ = Describe("Crawler", func() {
	Describe("#BuildSearchUrl", func() {
		Context("given the search string is valid", func() {
			It("returns the correct URL", func() {
				url, err := BuildSearchUrl("Ruby")
				if err != nil {
					Fail("Building search URL failed: " + err.Error())
				}

				Expect(url).To(Equal("https://www.google.com/search?hl=en&lr=lang_en&q=Ruby"))
			})
		})

		Context("given the search string is empty", func() {
			url, err := BuildSearchUrl("")
			if err != nil {
				Fail("Building search URL failed: " + err.Error())
			}

			Expect(url).To(Equal("https://www.google.com/search?hl=en&lr=lang_en&q="))
		})
	})
})

package crawler_test

import (
	"fmt"
	"net/url"

	. "google-scraper/services/crawler"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
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

		Context("given the search string is Thai string", func() {
			It("returns the correct URL", func() {
				queryString := "ทับทิม"
				searchUrl, err := BuildSearchUrl(queryString)
				if err != nil {
					Fail("Building search URL failed: " + err.Error())
				}

				encodedQueryString := url.QueryEscape(queryString)
				url := fmt.Sprintf("https://www.google.com/search?hl=en&lr=lang_en&q=%s", encodedQueryString)

				Expect(searchUrl).To(Equal(url))
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

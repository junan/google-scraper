package crawler_test

import (
	"io/ioutil"
	"strconv"
	"errors"

	. "google-scraper/services/crawler"
	. "google-scraper/helpers"
	. "google-scraper/constants"

	"github.com/jarcoal/httpmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Http", func() {
	Describe("#randomUserAgent", func() {
		It("returns with random user agent platform", func() {
			userAgent := RandomUserAgent()

			Expect(userAgent).To(MatchRegexp(`(Macintosh|Windows NT|Linux)`))
		})

		It("returns with random user agent browser", func() {
			userAgent := RandomUserAgent()

			Expect(userAgent).To(MatchRegexp(`(Firefox|Chrome)`))
		})
	})

	Describe("#GenerateRandomNumber", func() {
		It("returns the random number between 0 and 1", func() {
			randomNumber := strconv.Itoa(GenerateRandomNumber())

			Expect(randomNumber).To(MatchRegexp(`(1|0)`))
		})
	})


	Describe("#GetRequest", func() {
		Context("given Google search returns success response", func() {
			It("returns no error", func() {
				searchString := "Buy domain"
				responseString := "success response"
				searchUrl, err := BuildSearchUrl(searchString, GoogleSearchBaseUrl)
				if err != nil {
					Fail("Building search url failed: " + err.Error())
				}

				httpmock.RegisterResponder("GET", searchUrl,
					httpmock.NewStringResponder(200, responseString))

				result := httpmock.NewStringResponse(200, responseString)
				byte, err := ioutil.ReadAll(result.Body)
				if err != nil {
					Fail("Reading response failed: " + err.Error())
				}

				response, err := GetRequest(searchUrl)

				Expect(byte).To(Equal(response))
				Expect(err).To(BeNil())
			})
		})

		Context("given Google search returns error response", func() {
			It("returns an error", func() {
				searchUrl, err := BuildSearchUrl("Buy domain", GoogleSearchBaseUrl)
				if err != nil {
					Fail("Building search url failed: " + err.Error())
				}

				httpmock.RegisterResponder("GET", searchUrl,
					httpmock.NewErrorResponder(errors.New("some-error")))

				response, err := GetRequest(searchUrl)

				Expect(err).NotTo(BeNil())
				Expect(response).To(BeNil())
			})
		})
	})
})

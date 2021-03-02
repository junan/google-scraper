package crawler_test

import (
	"errors"
	"io/ioutil"

	. "google-scraper/constants"
	. "google-scraper/services/crawler"

	"github.com/jarcoal/httpmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Http", func() {
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
				body, err := ioutil.ReadAll(result.Body)
				if err != nil {
					Fail("Reading response failed: " + err.Error())
				}

				response, err := GetRequest(searchUrl)

				Expect(body).To(Equal(response))
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

				Expect(err.Error()).To(Equal( "Get \"https://www.google.com/search?hl=en&lr=lang_en&q=Buy+domain\": some-error"))
				Expect(response).To(BeNil())
			})
		})
	})
})

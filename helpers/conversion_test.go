package helpers_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"google-scraper/helpers"
)

var _ = Describe("Conversion", func() {
	Describe("#StringToInt", func() {
		Context("given a integer string", func() {
			It("returns the integer string as integer", func() {
				result, err := helpers.StringToInt("10")
				if err != nil {
					Fail("Conversion failed: " + err.Error())
				}
				Expect(result).To(BeNumerically("==", 10))
			})
		})

		Context("given the string is NOT integer string", func() {
			It("returns an err", func() {
				_, err := helpers.StringToInt("Hi")

				Expect(err).NotTo(BeNil())
			})
		})
	})
})

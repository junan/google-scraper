package helpers_test

import (
	"time"

	. "google-scraper/helpers"
	_ "google-scraper/initializers"
	. "google-scraper/tests"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("View", func() {
	Describe("#DisplayStatus", func() {
		Context("given the keyword search has not completed", func() {
			It("returns No", func() {
				user := FabricateUser("John", "john@example.com", "secret")
				keyword := FabricateKeyword("Buy domain", false, &user)

				Expect(DisplayStatus(keyword)).To(Equal("No"))
			})
		})

		Context("given the keyword search has completed", func() {
			It("returns Yes", func() {
				user := FabricateUser("John", "john@example.com", "secret")
				keyword := FabricateKeyword("Buy domain", true, &user)

				completed := DisplayStatus(keyword)
				Expect(completed).To(Equal("Yes"))
			})
		})
	})

	Describe("#DisplayFormattedCreatedDate", func() {
		It("returns formatted date", func() {
			currentTime := time.Now().Local()
			timeFormat := "January 2, 2006"
			user := FabricateUser("John", "john@example.com", "secret")
			keyword := FabricateKeyword("Buy domain", true, &user)

			formattedDate := DisplayFormattedCreatedDate(keyword)

			Expect(formattedDate).To(Equal(currentTime.Format(timeFormat)))
		})
	})

	AfterEach(func() {
		TruncateTables("users", "keywords")
	})
})

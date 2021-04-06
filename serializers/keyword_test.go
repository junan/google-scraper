package serializers_test

import (
	_ "google-scraper/initializers"
	"google-scraper/models"
	"google-scraper/serializers"
	. "google-scraper/tests"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("KeywordSerializer", func() {
	Describe("#Data", func() {
		It("returns correct data", func() {
			user := FabricateUser("John", "john@example.com", "secret")
			keyword1 := FabricateKeyword("Buy one domain", false, &user)
			keyword2 := FabricateKeyword("Buy two domain", false, &user)

			querySeterKeywords := models.GetQuerySeterKeywords(&user, "")
			keywords, err := models.GetPaginatedKeywords(querySeterKeywords, 0, 5)
			if err != nil {
				Fail("Getting paginated keywords failed: " + err.Error())
			}

			keywordsSerializer := serializers.KeywordList{
				Keywords: keywords,
			}

			data := keywordsSerializer.Data()

			Expect(len(data)).To(Equal(2))
			Expect(data[0].Id).To(Equal(keyword2.Id))
			Expect(data[1].Id).To(Equal(keyword1.Id))
		})
	})

	AfterEach(func() {
		TruncateTables("users", "keywords")
	})
})

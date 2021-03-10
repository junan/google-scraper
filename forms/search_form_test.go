package forms_test

import (
	"google-scraper/forms"
	_ "google-scraper/initializers"
	. "google-scraper/tests"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SearchForm", func() {
	Describe("#PerformSearch", func() {
		Context("given the search attributes are valid", func() {
			It("does NOT generate any error", func() {
				user := FabricateUser("John", "john@example.comn", "secret")
				validCsvFilePath := AppRootDir() + "/tests/fixtures/shared/valid_keywords.csv"
				file, header, err := GetFormFileData(validCsvFilePath)
				if err != nil {
					Fail("Getting form file data failed: " + err.Error())
				}

				mockResponseFilePath := AppRootDir() + "/tests/fixtures/services/crawler/valid_get_response.html"
				MockCrawling(mockResponseFilePath)

				err = forms.PerformSearch(file, header, &user)

				Expect(err).To(BeNil())
			})
		})

		Context("given the search attributes are INVALID", func() {
			Context("given the file is nil", func() {
				It("returns an error", func() {
					user := FabricateUser("John", "john@example.comn", "secret")

					err := forms.PerformSearch(nil, nil, &user)

					Expect(err.Error()).To(Equal("File can't be blank."))
				})
			})

			Context("given the file keywords are empty", func() {
				It("returns an error", func() {
					user := FabricateUser("John", "john@example.comn", "secret")
					filePath := AppRootDir() + "/tests/fixtures/shared/empty_keyword.csv"
					file, header, err := GetFormFileData(filePath)
					if err != nil {
						Fail("Getting form file data failed: " + err.Error())
					}

					err = forms.PerformSearch(file, header, &user)

					Expect(err.Error()).To(Equal("Keywords count can't be more than 1000 or less than 1."))
				})
			})

			Context("given the CSV file is wrong formatted", func() {
				It("returns an error", func() {
					user := FabricateUser("John", "john@example.comn", "secret")
					filePath := AppRootDir() + "/tests/fixtures/shared/invalid_keyword.csv"
					file, header, err := GetFormFileData(filePath)
					if err != nil {
						Fail("Getting form file data failed: " + err.Error())
					}

					err = forms.PerformSearch(file, header, &user)

					Expect(err.Error()).To(Equal("CSV contents are not in correct format."))
				})
			})

			Context("given the file is an image", func() {
				It("returns an error", func() {
					user := FabricateUser("John", "john@example.comn", "secret")
					filePath := AppRootDir() + "/tests/fixtures/shared/test.jpeg"
					file, header, err := GetFormFileData(filePath)
					if err != nil {
						Fail("Getting form file data failed: " + err.Error())
					}

					err = forms.PerformSearch(file, header, &user)

					Expect(err.Error()).To(Equal("Please upload the file in CSV format."))
				})
			})

			Context("given the file size is more than 5 megabytes", func() {
				It("returns an error", func() {
					user := FabricateUser("John", "john@example.comn", "secret")
					filePath := AppRootDir() + "/tests/fixtures/shared/big_file.pdf"
					file, header, err := GetFormFileData(filePath)
					if err != nil {
						Fail("Getting form file data failed: " + err.Error())
					}

					err = forms.PerformSearch(file, header, &user)

					Expect(err.Error()).To(Equal("File size can't be more than 5 MB."))
				})
			})
		})
	})

	AfterEach(func() {
		TruncateTables("users", "keywords", "search_results")
	})
})

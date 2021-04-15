package forms_test

import (
	"google-scraper/forms"
	_ "google-scraper/initializers"
	. "google-scraper/tests"

	"github.com/beego/beego/v2/core/validation"
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

			Context("given the CSV file is wrongly formatted", func() {
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

			Context("given the file is of an INVALID type", func() {
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

	Describe("#Valid", func() {
		Context("given the search attributes are valid", func() {
			It("does NOT add error to validation", func() {
				validCsvFilePath := AppRootDir() + "/tests/fixtures/shared/valid_keywords.csv"
				file, header, err := GetFormFileData(validCsvFilePath)
				if err != nil {
					Fail("Getting form file data failed: " + err.Error())
				}

				mockResponseFilePath := AppRootDir() + "/tests/fixtures/services/crawler/valid_get_response.html"
				MockCrawling(mockResponseFilePath)

				csv := forms.CSV{
					File:   file,
					Header: header,
					Size:   2,
				}

				formValidation := validation.Validation{}
				csv.Valid(&formValidation)

				Expect(len(formValidation.Errors)).To(BeZero())
			})
		})

		Context("given the search attributes are INVALID", func() {
			Context("given NO file", func() {
				It("adds an error to validation", func() {
					csv := forms.CSV{
						File:   nil,
						Header: nil,
						Size:   0,
					}

					formValidation := validation.Validation{}
					csv.Valid(&formValidation)

					Expect(len(formValidation.Errors)).To(Equal(1))
					Expect(formValidation.Errors[0].Message).To(Equal("File can't be blank."))
				})
			})

			Context("given the file keywords are empty", func() {
				It("adds an error to validation", func() {
					filePath := AppRootDir() + "/tests/fixtures/shared/empty_keyword.csv"
					file, header, err := GetFormFileData(filePath)
					if err != nil {
						Fail("Getting form file data failed: " + err.Error())
					}

					csv := forms.CSV{
						File:   file,
						Header: header,
						Size:   0,
					}

					formValidation := validation.Validation{}
					csv.Valid(&formValidation)

					Expect(len(formValidation.Errors)).To(Equal(1))
					Expect(formValidation.Errors[0].Message).To(Equal("Keywords count can't be more than 1000 or less than 1."))
				})
			})

			Context("given the CSV file is wrongly formatted", func() {
				It("adds an error to validation", func() {
					filePath := AppRootDir() + "/tests/fixtures/shared/invalid_keyword.csv"
					file, header, err := GetFormFileData(filePath)
					if err != nil {
						Fail("Getting form file data failed: " + err.Error())
					}

					csv := forms.CSV{
						File:   file,
						Header: header,
						Size:   0,
					}

					formValidation := validation.Validation{}
					csv.Valid(&formValidation)

					Expect(len(formValidation.Errors)).To(Equal(1))
					Expect(formValidation.Errors[0].Message).To(Equal("CSV contents are not in correct format."))
				})
			})

			Context("given the file is of an INVALID type", func() {
				It("adds an error to validation", func() {
					filePath := AppRootDir() + "/tests/fixtures/shared/test.jpeg"
					file, header, err := GetFormFileData(filePath)
					if err != nil {
						Fail("Getting form file data failed: " + err.Error())
					}

					csv := forms.CSV{
						File:   file,
						Header: header,
						Size:   0,
					}

					formValidation := validation.Validation{}
					csv.Valid(&formValidation)

					Expect(len(formValidation.Errors)).To(Equal(1))
					Expect(formValidation.Errors[0].Message).To(Equal("Please upload the file in CSV format."))
				})
			})

			Context("given the file size is more than 5 megabytes", func() {
				It("adds an error to validation", func() {
					filePath := AppRootDir() + "/tests/fixtures/shared/big_file.pdf"
					file, header, err := GetFormFileData(filePath)
					if err != nil {
						Fail("Getting form file data failed: " + err.Error())
					}

					csv := forms.CSV{
						File:   file,
						Header: header,
						Size:   0,
					}

					formValidation := validation.Validation{}
					csv.Valid(&formValidation)

					Expect(len(formValidation.Errors)).To(Equal(1))
					Expect(formValidation.Errors[0].Message).To(Equal("Please upload the file in CSV format."))
				})
			})
		})
	})

	AfterEach(func() {
		TruncateTables("users", "keywords", "search_results")
	})
})

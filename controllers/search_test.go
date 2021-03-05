package controllers_test

import (
	"bytes"
	"mime/multipart"

	. "google-scraper/tests/fabricators"
	. "google-scraper/tests/testing_helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SearchController", func() {
	AfterEach(func() {
		TruncateTable("users")
		TruncateTable("keywords")
		TruncateTable("search_results")
	})

	Describe("POST", func() {
		Context("given the user is an authenticated user", func() {
			Context("given the user is uploaded a valid CSV file", func() {
				It("sets flash success message", func() {
					email := "john@example.com"
					password := "secret"
					user := FabricateUser("John", email, password)
					//params := url.Values{
					//	"email":    {email},
					//	"password": {password},
					//}

					body := new(bytes.Buffer)
					writer := multipart.NewWriter(body)
					writer.WriteField("bu", "HFL")
					writer.WriteField("wk", "10")
					part, _ := writer.CreateFormFile("file", "file.csv")
					part.Write([]byte(`sample`))
					writer.Close() // <<< important part

					response := MakeAuthenticatedRequest("POST", "/search", body, writer.FormDataContentType(), &user)

					flash := GetFlash(response.Cookies())

					Expect(flash.Data["success"]).To(Equal("Your csv file has been uploaded successfully"))
				})

				//It("redirects to the root path", func() {
				//	email := "john@example.com"
				//	password := "secret"
				//	user := FabricateUser("John", email, password)
				//	registrationForm := url.Values{
				//		"name":     {"John"},
				//		"email":    {email},
				//		"password": {password},
				//	}
				//	body := strings.NewReader(registrationForm.Encode())
				//	response := MakeAuthenticatedRequest("POST", "/register", body, &user)
				//	path := GetUrlPath(response)
				//
				//	Expect(path).To(Equal("/"))
				//})
			})
		})
	})
})

package controllers_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "google-scraper/tests/testing_helpers"
	"net/http"
)

var _ = Describe("RegistrationController", func() {
	AfterEach(func() {
		TruncateTable("users")
	})

	Describe("GET", func() {
		It("returns 200 status code", func() {
			response := MakeRequest("GET", "/register", nil)

			Expect(response.Code).To(Equal(http.StatusOK))
		})
	})

	//Describe("POST", func() {
	//	Context("given valid params", func() {
	//		It("returns status FOUND", func() {
	//			form := url.Values{
	//				"email":                 {"hoang.mirs@gmail.com"},
	//				"password":              {"123456"},
	//				"password_confirmation": {"123456"},
	//			}
	//			body := strings.NewReader(form.Encode())
	//			response := MakeRequest("POST", "/register", body)
	//
	//			Expect(response.Code).To(Equal(http.StatusFound))
	//		})
	//	})
	//
	//	Context("given invalid params", func() {
	//		It("returns status OK", func() {
	//			form := url.Values{
	//				"email":                 {""},
	//				"password":              {""},
	//				"password_confirmation": {""},
	//			}
	//			body := strings.NewReader(form.Encode())
	//			response := MakeRequest("POST", "/register", body)
	//
	//			Expect(response.Code).To(Equal(http.StatusOK))
	//			This conversation was marked as resolved by olivierobert
	//		})
	//
	//		It("returns error flash message", func() {
	//			form := url.Values{
	//				"email":                 {""},
	//				"password":              {""},
	//				"password_confirmation": {""},
	//			}
	//			body := strings.NewReader(form.Encode())
	//			response := MakeRequest("POST", "/register", body)
	//
	//			flashMessage := GetFlash(response.Result().Cookies())
	//
	//			Expect(flashMessage.Data["error"]).To(Equal("Email can not be empty"))
	//		})
	//	})
	//})
})

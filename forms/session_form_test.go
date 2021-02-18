package forms_test

import (
	"google-scraper/forms"
	_ "google-scraper/initializers"
	. "google-scraper/tests/fabricators"
	. "google-scraper/tests/testing_helpers"

	"github.com/beego/beego/v2/core/validation"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SessionForm", func() {
	Describe("#Valid", func() {
		Context("given the session params are valid", func() {
			It("does NOT generate any errors", func() {
				email := "john@example.com"
				password := "secret"
				FabricateUser("John", email, password)
				form := forms.SessionForm{
					Email:    email,
					Password: password,
				}

				formValidation := validation.Validation{}
				form.Valid(&formValidation)

				Expect(len(formValidation.Errors)).To(BeZero())
			})
		})

		Context("given the session params are INVALID", func() {
			Context("given the user email is INVALID", func() {
				It("adds an error to the email field", func() {
					email := "john@example.com"
					password := "secret"
					FabricateUser("John", email, password)
					form := forms.SessionForm{
						Email:    "invalid_email",
						Password: password,
					}

					formValidation := validation.Validation{}
					form.Valid(&formValidation)

					Expect(len(formValidation.Errors)).To(Equal(1))
					Expect(formValidation.Errors[0].Key).To(Equal("Email"))
					Expect(formValidation.Errors[0].Message).To(Equal("Incorrect email or password"))
				})
			})

			Context("given the user email does NOT exist", func() {
				It("adds an error to the email field", func() {
					form := forms.SessionForm{
						Email:    "non_existing_email@example.com",
						Password: "secret",
					}

					formValidation := validation.Validation{}
					form.Valid(&formValidation)

					Expect(len(formValidation.Errors)).To(Equal(1))
					Expect(formValidation.Errors[0].Key).To(Equal("Email"))
					Expect(formValidation.Errors[0].Message).To(Equal("Incorrect email or password"))
				})
			})

			Context("given the user password is INVALID", func() {
				It("adds an error to the password field", func() {
					email := "john@example.com"
					password := "secret"
					FabricateUser("John", email, password)

					form := forms.SessionForm{
						Email:    email,
						Password: "wrong-password",
					}

					formValidation := validation.Validation{}
					form.Valid(&formValidation)

					Expect(len(formValidation.Errors)).To(Equal(1))
					Expect(formValidation.Errors[0].Key).To(Equal("Password"))
					Expect(formValidation.Errors[0].Message).To(Equal("Incorrect email or password"))
				})
			})
		})
	})

	Describe("#Authenticate", func() {
		Context("given the session params are valid", func() {
			It("returns the user with no error", func() {
				email := "john@example.com"
				password := "secret"
				user := FabricateUser("John", email, password)
				form := forms.SessionForm{
					Email:    email,
					Password: password,
				}

				currentUser, err := form.Authenticate()

				Expect(err).To(BeNil())
				Expect(currentUser.Id).To(Equal(user.Id))
			})
		})

		Context("given the session params are INVALID", func() {
			Context("given the user email is INVALID", func() {
				It("returns a generic error message", func() {
					form := forms.SessionForm{
						Email:    "invalid-email",
						Password: "secret",
					}

					user, err := form.Authenticate()

					Expect(err.Error()).To(Equal("Incorrect email or password"))
					Expect(user).To(BeNil())
				})
			})

			Context("given the user email is empty", func() {
				It("returns a generic error message", func() {
					form := forms.SessionForm{
						Email:    "",
						Password: "secret",
					}

					user, err := form.Authenticate()

					Expect(err.Error()).To(Equal("Incorrect email or password"))
					Expect(user).To(BeNil())
				})
			})

			Context("given the user password is wrong", func() {
				It("returns a generic error message", func() {
					email := "john@example.com"
					password := "secret"
					FabricateUser("John", email, password)

					form := forms.SessionForm{
						Email:    email,
						Password: "wrong-password",
					}
					user, err := form.Authenticate()

					Expect(err.Error()).To(Equal("Incorrect email or password"))
					Expect(user).To(BeNil())
				})
			})

			Context("given the user password is blank", func() {
				It("returns a generic error message", func() {
					email := "john@example.com"
					password := "secret"
					FabricateUser("John", email, password)

					form := forms.SessionForm{
						Email:    email,
						Password: "",
					}
					user, err := form.Authenticate()

					Expect(err.Error()).To(Equal("Incorrect email or password"))
					Expect(user).To(BeNil())
				})
			})
		})
	})

	AfterEach(func() {
		TruncateTable("users")
	})
})

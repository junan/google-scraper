package forms_test

import (
	"google-scraper/forms"
	_ "google-scraper/initializers"
	. "google-scraper/tests"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RegistrationForm", func() {
	Describe("Save", func() {
		Context("given valid attributes", func() {
			It("creates a new user and returns nil error", func() {
				registrationForm := forms.RegistrationForm{
					Name:     "John",
					Email:    "john@example.com",
					Password: "secret",
				}
				user, err := registrationForm.Save()

				Expect(user.Id).NotTo(BeNil())
				Expect(err).To(BeNil())
			})
		})

		Context("given invalid attributes", func() {
			Context("given a blank email", func() {
				It("returns the correct error message and does NOT create new user", func() {
					registrationForm := forms.RegistrationForm{
						Name:     "John",
						Email:    "",
						Password: "secret",
					}
					user, err := registrationForm.Save()

					Expect(user).To(BeNil())
					Expect(err.Error()).To(Equal("Email can not be empty"))
				})
			})

			Context("given an existing email", func() {
				It("returns the correct error message and does NOT create new user", func() {
					email := "john@example.com"
					FabricateUser("John", email, "secret")

					registrationForm := forms.RegistrationForm{
						Name:     "John",
						Email:    email,
						Password: "secret",
					}

					user, err := registrationForm.Save()

					Expect(user).To(BeNil())
					Expect(err.Error()).To(Equal("Email already exists"))
				})
			})

			Context("given an invalid email", func() {
				It("returns the correct error message and does NOT create new user", func() {
					registrationForm := forms.RegistrationForm{
						Name:     "John",
						Email:    "invalid-email",
						Password: "secret",
					}
					user, err := registrationForm.Save()

					Expect(user).To(BeNil())
					Expect(err.Error()).To(Equal("Email must be a valid email address"))
				})
			})

			Context("given a blank name", func() {
				It("returns the correct error message and does NOT create new user", func() {
					registrationForm := forms.RegistrationForm{
						Name:     "",
						Email:    "john@example.com",
						Password: "secret",
					}
					user, err := registrationForm.Save()

					Expect(user).To(BeNil())
					Expect(err.Error()).To(Equal("Name can not be empty"))
				})
			})

			Context("given a blank password", func() {
				It("returns the correct error message and does NOT create new user", func() {
					registrationForm := forms.RegistrationForm{
						Name:     "John",
						Email:    "john@example.com",
						Password: "",
					}
					user, err := registrationForm.Save()

					Expect(user).To(BeNil())
					Expect(err.Error()).To(Equal("Password can not be empty"))
				})
			})

			Context("given a short password by three character", func() {
				It("returns an error message and does NOT create a new user", func() {
					registrationForm := forms.RegistrationForm{
						Name:     "John",
						Email:    "john@example.com",
						Password: "abc",
					}
					user, err := registrationForm.Save()

					Expect(user).To(BeNil())
					Expect(err.Error()).To(Equal("Password minimum size is 6"))
				})
			})
		})
	})

	AfterEach(func() {
		TruncateTables("users")
	})
})

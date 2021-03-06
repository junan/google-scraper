package forms

import (
	"google-scraper/helpers"
	"google-scraper/models"
	"google-scraper/constants"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/validation"
)

type RegistrationForm struct {
	Name     string `form:"name" valid:"Required;"`
	Email    string `form:"email" valid:"Required; Email; MaxSize(100)"`
	Password string `form:"password" valid:"Required; MinSize(6)"`
}

func (registrationForm *RegistrationForm) Valid(v *validation.Validation) {
	user := models.User{
		Email: registrationForm.Email,
	}

	if user.IsExistingUser(){
		_ = v.SetError("Email", "Email already exists")
	}
}

func (registrationForm *RegistrationForm) Save() (*models.User, error) {
	validation := validation.Validation{}

	success, err := validation.Valid(registrationForm)

	if err != nil {
		logs.Error(constants.GeneralValidationFailedLogMessage, err)
	}

	if !success {
		for _, err := range validation.Errors {
			return nil, err
		}
	}

	hashedPassword, err := helpers.HashPassword(registrationForm.Password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name:           registrationForm.Name,
		Email:          registrationForm.Email,
		HashedPassword: hashedPassword,
	}

	_, err = models.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, err
}

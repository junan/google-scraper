package forms

import (
	"fmt"

	"google-scraper/constants"
	"google-scraper/helpers"
	"google-scraper/models"

	"github.com/beego/beego/v2/adapter/logs"
	"github.com/beego/beego/v2/core/validation"
)

const (
	sessionFormErrorMessage = "Incorrect email or password"
	sessionFormLogMessage = "Setting error on validation failed: "
)

type SessionForm struct {
	Email    string `form:"email" valid:"Email; Required"`
	Password string `form:"password" valid:"Required;"`
}

var currentUser *models.User

func (sessionForm *SessionForm) Valid(v *validation.Validation) {
	user, err := models.FindUserByEmail(sessionForm.Email)
	if err != nil {
		err = v.SetError("Email", sessionFormErrorMessage)
		if err == nil {
			logs.Error(sessionFormLogMessage, err)
		}
	} else {
		err = helpers.VerifyPasswordHash(sessionForm.Password, user.HashedPassword)
		if err != nil {
			err = v.SetError("Password", sessionFormErrorMessage)
			if err == nil {
				logs.Error(sessionFormLogMessage, err)
			}
		} else {
			currentUser = user
		}
	}
}

func (sessionForm *SessionForm) Authenticate() (*models.User, error) {
	validation := validation.Validation{}
	success, err := validation.Valid(sessionForm)

	if err != nil {
		logs.Error(constants.GeneralValidationFailedLogMessage, err)
	}

	if !success {
		return nil, fmt.Errorf(sessionFormErrorMessage)
	}

	return currentUser, err
}

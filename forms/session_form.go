package forms

import (
	"google-scraper/constants"
	"google-scraper/helpers"
	"google-scraper/models"

	"github.com/beego/beego/v2/adapter/logs"
	"github.com/beego/beego/v2/core/validation"
)

type SessionForm struct {
	Email    string `form:"email" valid:"Email; Required"`
	Password string `form:"password" valid:"Required;"`
}

var currentUser *models.User

func (sessionForm *SessionForm) Valid(v *validation.Validation) {
	user, err := models.FindUserByEmail(sessionForm.Email)
	if err != nil {
		err := v.SetError("Email", constants.SessionValidationErrorMessage)
		if err == nil {
			logs.Error(constants.SessionFormValidationFailedLogMessage, err)
		}
	} else {
		err = helpers.CheckPasswordHash(sessionForm.Password, user.HashedPassword)
		if err != nil {
			err := v.SetError("Password", constants.SessionValidationErrorMessage)
			if err == nil {
				logs.Error(constants.SessionFormValidationFailedLogMessage, err)
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
		for _, err := range validation.Errors {
			return nil, err
		}
	}

	return currentUser, err
}

package forms

import (
	"google-scraper/helpers"
	"google-scraper/models"

	"github.com/beego/beego/v2/adapter/logs"
	"github.com/beego/beego/v2/core/validation"
)

type SessionForm struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

var currentUser *models.User

func (sessionForm *SessionForm) Valid(v *validation.Validation) {
	errMessage := "Incorrect email or password"
	logMessage := "Failed to set error on validation: "

	user, err := models.FindUserByEmail(sessionForm.Email)
	if err != nil {
		err := v.SetError("Email", errMessage)
		if err == nil {
			logs.Error(logMessage, err)
		}
	} else {
		err = helpers.CheckPasswordHash(sessionForm.Password, user.HashedPassword)
		if err != nil {
			err := v.SetError("Password", errMessage)
			if err == nil {
				logs.Error(logMessage, err)
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
		logs.Error("Validation error:", err)
	}

	if !success {
		for _, err := range validation.Errors {
			return nil, err
		}
	}

	return currentUser, err
}

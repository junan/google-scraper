package forms

import (
	"github.com/beego/beego/v2/adapter/logs"
	"github.com/beego/beego/v2/core/validation"
	"google-scraper/helpers"
	"google-scraper/models"
)

type SessionForm struct {
	Email    string `form:"email" valid:"Required; MaxSize(100)"`
	Password string `form:"password" valid:"Required; MinSize(6)"`
}

var currentUser *models.User

func (sessionForm *SessionForm) Valid(v *validation.Validation) {
	errMessage := "Incorrect email or password"
	LogMessage := "Failed to set error on validation: "
	user, err := models.FindUserByEmail(sessionForm.Email)
	if err != nil {
		err := v.SetError("Email", errMessage)
		if err == nil {
			logs.Error(LogMessage, err)
		}
	} else {
		err = helpers.CheckPasswordHash(sessionForm.Password, user.HashedPassword)
		if err != nil {
			err := v.SetError("Password", errMessage)
			if err == nil {
				logs.Error(LogMessage, err)
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

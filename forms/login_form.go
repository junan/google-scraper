package forms

import (
	"github.com/beego/beego/v2/adapter/logs"
	"github.com/beego/beego/v2/adapter/validation"
	"github.com/beego/beego/v2/client/orm"
	"google-scraper/helpers"
	"google-scraper/models"
)

type SessionForm struct {
	Email    string `form:"email" valid:"Required; MaxSize(100)"`
	Password string `form:"password" valid:"Required; MinSize(6)"`
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

	user := models.User{
		Email: sessionForm.Email,
	}

	o := orm.NewOrm()
	err = o.Read(&user, "Email")

	_, err = helpers.CheckPasswordHash(sessionForm.Password, user.HashedPassword)
	if err != nil {
		return &user, err
	}

	return &user, err
}

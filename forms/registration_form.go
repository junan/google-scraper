package forms

import (
	"google-scraper/helpers"
	"google-scraper/models"

	"github.com/beego/beego/v2/client/orm"
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

	o := orm.NewOrm()
	_ = o.Read(&user, "Email")

	if user.Id != 0 {
		_ = v.SetError("Email", "Email already exists")
	}
}

func (registrationForm *RegistrationForm) Save() (*models.User, error) {
	validation := validation.Validation{}

	success, err := validation.Valid(registrationForm)

	if err != nil {
		logs.Error("Validation error:", err)
	}

	if !success {
		for _, err := range validation.Errors {
			return nil, err
		}
	}

	user := &models.User{
		Name:  registrationForm.Name,
		Email: registrationForm.Email,
	}
	user.HashedPassword = helpers.HashPassword(registrationForm.Password)

	_, err = models.CreateUser(user)

	return user, err
}

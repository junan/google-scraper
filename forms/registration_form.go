package forms

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/validation"
	"google-scraper/helpers"
	"google-scraper/models"
)

type RegistrationForm struct {
	Name     string `form:"name" valid:"Required;"`
	Email    string `form:"email" valid:"Required; Email; MaxSize(100)"`
	Password string `form:"password" valid:"Required; MinSize(6)"`
}

func (registrationForm *RegistrationForm) Valid(v *validation.Validation) {

	fmt.Println("Junan")

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
	fmt.Println("validation scope")
	validation := validation.Validation{}

	success, err := validation.Valid(registrationForm)
	if err != nil {
		logs.Error("Validate err:", err)
	}

	if !success {
		for _, err := range validation.Errors {
			return nil, err
		}
	}

	user := &models.User{
		Email: registrationForm.Email,
	}

	user.HashedPassword = helpers.HashPassword(registrationForm.Password)

	o := orm.NewOrm()
	_, err = o.Insert(user)

	return user, err
}

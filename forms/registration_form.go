package forms

import (
	"google-scraper/helpers"
	"google-scraper/models"
	"github.com/beego/beego/v2/client/orm"
)

type RegistrationForm struct {
	Name     string `form:"name" valid:"Required;"`
	Email    string `form:"email" valid:"Required; Email; MaxSize(100)"`
	Password string `form:"password" valid:"Required; MinSize(6)"`
}

func (form RegistrationForm) Save() error {
	var err error
	user := models.User{}
	user.Email = form.Email
	user.Name = form.Name
	user.HashedPassword = helpers.HashPassword(form.Password)

	o := orm.NewOrm()
	_, err = o.Insert(&user)

	return err
}

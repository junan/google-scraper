package forms

import (
	"errors"
	"google-scraper/models"

	"github.com/beego/beego/v2/core/validation"
)

type SearchForm struct {
	File string `form:"file" valid:"Required;"`
}

func (searchForm *SearchForm) Valid(v *validation.Validation) {
	//user := models.User{
	//	Email: searchForm.File,
	//}
	//
	//if user.IsExistingUser(){
	//	_ = v.SetError("Email", "Email already exists")
	//}
}

func (searchForm *SearchForm) Save(user *models.User) (*models.Keyword, error) {
	//validation := validation.Validation{}

	//success, err := validation.Valid(searchForm)

	//if err != nil {
	//	logs.Error(constants.GeneralValidationFailedLogMessage, err)
	//}
	//
	//if !success {
	//	for _, err := range validation.Errors {
	//		return nil, err
	//	}
	//}

	//
	search := &models.Keyword{
		Name: "Junan",
		User: user,
	}

	err := errors.New("math: square root of negative number")

	return search, err
}

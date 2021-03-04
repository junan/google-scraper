package forms

import (
	"encoding/csv"
	"fmt"
	"mime/multipart"
	"reflect"

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

func SearchProcess(file multipart.File) ([][]string, error) {
	records, err := readData(file)

	return records, err
}

func readData(file multipart.File) ([][]string, error) {
	//fileExt := path.Ext(head.Filename)
	//if fileExt != ".jpg" && fileExt != ".png" && fileExt != ".jpeg"{
	//	beego.Info ("The uploaded image format is incorrect, please add it again!")
	//	this.TplName = "add.html"
	//	return
	//}

	r := csv.NewReader(file)

	// skip first line
	_, err := r.Read()
	if err != nil {
		return [][]string{}, err
	}

	records, err := r.ReadAll()
	tj := reflect.TypeOf(records)
	fmt.Println(tj)
	if err != nil {
		return [][]string{}, err
	}

	return records, nil
}

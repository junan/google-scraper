package forms

import (
	"encoding/csv"
	"mime/multipart"
	"path"

	"github.com/beego/beego/v2/core/validation"
)

var keywords [][]string
var allowedExtensionMap = map[string]bool{
	".csv": true,
}

type CSV struct {
	File   multipart.File
	Header *multipart.FileHeader
}

func (csv *CSV) Valid(v *validation.Validation) {
	// Verifying file is not empty
	if csv.File == nil {
		_ = v.SetError("File", "Uploaded file can't be blank")
	}

	// Verifying uploaded file is in CSV format
	extension := path.Ext(csv.Header.Filename)
	_, ok := allowedExtensionMap[extension]
	if !ok {
		_ = v.SetError("File", "Uploaded file can't be blank")
	}

	// Verifying csv format
	_, err := readKeywords(csv.File)
	if err != nil {
		_ = v.SetError("File", err.Error())
	}
}

func SearchProcess(file multipart.File, header *multipart.FileHeader) (errs []error) {
	csvFile := CSV{File: file, Header: header}
	validation := validation.Validation{}
	success, err := validation.Valid(csvFile)

	if err != nil {
		return errs
	}

	if !success {
		for _, err := range validation.Errors {
			errs = append(errs, err)
		}
	}

	// Do the cron jobs

	return errs
}

func readKeywords(file multipart.File) ([][]string, error) {
	r := csv.NewReader(file)

	// skip csv header
	_, err := r.Read()
	if err != nil {
		return [][]string{}, err
	}

	keywords, err := r.ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return keywords, nil
}

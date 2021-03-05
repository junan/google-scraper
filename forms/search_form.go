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
	Size   int64
}

func (csv *CSV) Valid(v *validation.Validation) {
	// Verifying file is not empty
	if csv.File == nil {
		_ = v.SetError("File", "File can't be blank")
		return
	}

	// Verifying file size is not more than 5MB
	if csv.Size > 5 {
		_ = v.SetError("Size", "File size can't be more than 5 MB")
		return
	}

	// Verifying uploaded file is in CSV format
	extension := path.Ext(csv.Header.Filename)
	_, ok := allowedExtensionMap[extension]
	if !ok {
		_ = v.SetError("File", "File should be in CSV format")
		return
	}

	// Verifying csv format
	_, err := readKeywords(csv.File)
	if err != nil {
		_ = v.SetError("File", firstCapitalise(err.Error()))
	}
}

func PerformSearch(file multipart.File, header *multipart.FileHeader) (err error) {
	csvFile := CSV{File: file, Header: header, Size: getSizeInMb(header)}
	validation := validation.Validation{}
	success, err := validation.Valid(&csvFile)

	if err != nil {
		return err
	}

	if !success {
		for _, err := range validation.Errors {
			return err
		}
	}

	// Do the cron jobs

	return err
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

func getSizeInMb(header *multipart.FileHeader) int64 {
	var size int64
	if header == nil {
		size = 0
	} else {
		size = header.Size / (1024 * 1024)
	}

	return size
}

package forms

import (
	"encoding/csv"
	"errors"
	"mime/multipart"
	"path"
)

var AllowExtensionMap = map[string]bool{
	".csv": true,
}

func SearchProcess(file multipart.File, header *multipart.FileHeader) []error {
	errs := validateCSVFile(file, header)
	if len(errs) > 0 {
		return errs
	}
	_, err := readKeywords(file)
	if err != nil {
		return []error{err}
	}

	// Do the cron jobs

	return []error{}
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

func validateCSVFile(file multipart.File, header *multipart.FileHeader) []error {
	// Verifying file is not empty
	if file == nil {
		err := errors.New("Uploaded file cant be empty")
		return []error{err}
	}

	// Verifying uploaded file is in CSV format
	extension := path.Ext(header.Filename)
	_, ok := AllowExtensionMap[extension]
	if !ok {
		err := errors.New("The uploaded file needs to be in CSV format")
		return []error{err}
	}

	return []error{}
}

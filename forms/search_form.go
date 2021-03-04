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

func SearchProcess(file multipart.File, header *multipart.FileHeader) ([]error) {
	errs := validateCSVFile(file, header)
	if len(errs) >= 0 {
		return errs
	}
	_, err := readData(file)
	// Do the cron jobs
	if err != nil {
		return []error{err}
	}

	return []error{}
}

func readData(file multipart.File) ([][]string, error) {
	r := csv.NewReader(file)

	// skip first line
	_, err := r.Read()
	if err != nil {
		return [][]string{}, err
	}

	records, err := r.ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return records, nil
}

func validateCSVFile(file multipart.File, header *multipart.FileHeader) []error {
	if file == nil {
		err := errors.New("You need to upload some file")
		return []error{err}
	}
	extension := path.Ext(header.Filename)
	_, ok := AllowExtensionMap[extension]
	if !ok {
		err := errors.New("The uploaded file needs to be in CSV format")
		return []error{err}
	}

	return []error{}
}

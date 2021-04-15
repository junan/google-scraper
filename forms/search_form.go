package forms

import (
	"encoding/csv"
	"mime/multipart"
	"path"

	"google-scraper/models"
	. "google-scraper/services/enqueueing"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/validation"
)

var keywordStrings [][]string
var CsvKeywordValidationCriteria = [...]string{"presence", "size", "extension", "format", "count"}
var CsvValidationMessageMapping = map[string]string{
	"presence":  "File can't be blank.",
	"size":      "File size can't be more than 5 MB.",
	"extension": "Please upload the file in CSV format.",
	"format":    "CSV contents are not in correct format.",
	"count":     "Keywords count can't be more than 1000 or less than 1.",
}
var allowedExtensionMap = map[string]bool{
	".csv": true,
}

type CSV struct {
	File   multipart.File
	Header *multipart.FileHeader
	Size   int64
}

func (csv *CSV) Valid(v *validation.Validation) {
	for _, criteria := range CsvKeywordValidationCriteria {
		success := validate(criteria, csv)
		if !success {
			_ = v.SetError("File", CsvValidationMessageMapping[criteria])
			break
		}
	}
}

func PerformSearch(file multipart.File, header *multipart.FileHeader, user *models.User) (err error) {
	var keywordIndex int64 = 0
	csvFile := CSV{File: file, Header: header, Size: getSizeInMb(header)}
	validation := validation.Validation{}
	success, err := validation.Valid(&csvFile)

	if err != nil {
		return err
	}

	if !success {
		for _, err := range validation.Errors {
			// Returning only first error, as there will be only one error
			return err
		}
	}

	for _, row := range keywordStrings {
		for _, name := range row {
			keyword, err := storeKeyword(name, user)
			if err == nil {
				job, err := EnqueueKeywordJob(keyword, keywordIndex)
				if err != nil {
					logs.Error("Adding keyword to queue failed: ", err)
				}
				logs.Info("Enqueued %v keyword to the %v", keyword.Name, job.Name)
				keywordIndex = keywordIndex + 1
			}
		}
	}

	return nil
}

func readKeywords(file multipart.File) ([][]string, error) {
	r := csv.NewReader(file)

	// skip csv header
	_, err := r.Read()
	if err != nil {
		return [][]string{}, err
	}

	keywordStrings, err = r.ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return keywordStrings, nil
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

func validate(criteria string, csv *CSV) bool {
	if criteria == "presence" {
		return validateFilePresence(csv)
	} else if criteria == "size" {
		return validateFileSize(csv)
	} else if criteria == "extension" {
		return validateFileExtension(csv)
	} else if criteria == "format" {
		return validateKeywordFormat(csv)
	} else if criteria == "count" {
		return validateKeywordCount(keywordStrings)
	}

	return true
}

func validateFilePresence(csv *CSV) bool {
	return csv.File != nil
}

func validateFileSize(csv *CSV) bool {
	return csv.Size < 5 // 5 = five megabyte
}

func validateFileExtension(csv *CSV) bool {
	extension := path.Ext(csv.Header.Filename)
	_, ok := allowedExtensionMap[extension]

	return ok
}

func validateKeywordFormat(csv *CSV) bool {
	_, err := readKeywords(csv.File)

	return err == nil
}

func validateKeywordCount(keywords [][]string) bool {
	row := len(keywords)
	switch row {
	case 0:
		return false
	case 1:
		return len(keywords[0]) <= 1000
	default:
		column := len(keywords[1])
		totalCount := row * column
		return totalCount <= 1000
	}
}

func storeKeyword(name string, user *models.User) (keyword *models.Keyword, err error) {
	keyword = &models.Keyword{
		Name: name,
		User: user,
	}

	_, err = models.CreateKeyword(keyword)
	if err != nil {
		return nil, err
	}

	return keyword, nil
}

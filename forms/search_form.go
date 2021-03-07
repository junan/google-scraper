package forms

import (
	"encoding/csv"
	"mime/multipart"
	"path"

	"google-scraper/models"
	. "google-scraper/services/crawler"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/validation"
)

var keywordStrings [][]string
var keywordIds []int64
var CsvKeywordValidationCriteria = [...]string{"presence", "size", "extension", "format", "count"}
var CSVValidationMessageMapping = map[string]string{
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
			v.SetError("File", CSVValidationMessageMapping[criteria])
			break
		}
	}
}

func PerformSearch(file multipart.File, header *multipart.FileHeader, user *models.User) (err error) {
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

	// TODO: This part will be processed in cron job, will be added some request delay technique and requeue the job on fails
	// Storing the search string in the Keyword model and creating SearchResult model with the crawled data using the keyword object
	for _, row := range keywordStrings {
		for _, name := range row {
			keyword, err := storeKeyword(name, user)
			if err == nil {
				Crawl(keyword)
			}
		}
	}

	return err
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

func storeKeyword(name string, user *models.User) (keyword *models.Keyword, err error) {
	keyword = &models.Keyword{
		Name: name,
		User: user,
	}

	id, err := models.CreateKeyword(keyword)
	if err != nil {
		logs.Error("Creating keyword failed: ", err)

	} else {
		result := append(keywordIds, id)
		keywordIds = result
	}

	return keyword, err
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
	if err != nil {
		return false
	}

	return true
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

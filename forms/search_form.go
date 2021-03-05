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
			keyword, err :=  storeKeyword(name, user)
			if err == nil {
				Crawl(keyword)
			}
		}
	}

	return err
}

func readKeywords(file multipart.File) ([][]string, error) {
	r := csv.NewReader(file)
	// To support semicolon separated csv and comments
	r.Comma = ';'
	r.Comment = '#'

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

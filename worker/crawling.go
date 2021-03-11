package worker

import (
	"google-scraper/models"
	. "google-scraper/services/crawler"

	"github.com/beego/beego/v2/core/logs"
)

type Crawling struct {
	CsvKeywords [][]string
	UserId int64
}

func (crawling *Crawling) PerformLater() error {
	for _, row := range crawling.CsvKeywords {
		for _, name := range row {
			keyword, err := storeKeyword(name, crawling.UserId)
			if err == nil {
				_, err = Crawl (keyword)
				if err != nil {
					logs.Error("Crawling failed: ", err)
				}
			}
		}
	}

	return nil
}

func storeKeyword(name string, userId int64) (keyword *models.Keyword, err error) {
	user, err := models.FindUserById(userId)
	if err != nil {
		return nil, err
	}

	keyword = &models.Keyword{
		Name: name,
		User: user,
	}

	_, err = models.CreateKeyword(keyword)
	if err != nil {
		return nil, err
	}

	return keyword, err
}

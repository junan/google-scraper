package helpers

import (
	"google-scraper/models"
)

const (
	layoutUS  = "January 2, 2006"
)

func DisplayStatus(k models.Keyword) string {
	if k.SearchCompleted {
		return "Yes"
	} else {
		return "No"
	}
}

func DisplayFormattedCreatedDate(k models.Keyword) string {
    return k.CreatedAt.Format(layoutUS)
}

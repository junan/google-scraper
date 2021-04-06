package tests

import (
	"google-scraper/services/oauth"

	"github.com/go-oauth2/oauth2/v4/models"
	. "github.com/onsi/ginkgo"
)

func FabricateOAuthClient(id string, secret string) *models.Client {
	client := &models.Client{
		ID:     id,
		Secret: secret,
	}

	err := oauth.ClientStore.Create(client)
	if err != nil {
		Fail("Creating client failed:" + err.Error())
	}

	return client
}

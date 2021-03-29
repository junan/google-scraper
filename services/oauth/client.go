package oauth

import (
	. "google-scraper/helpers"

	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/google/uuid"
)

func GenerateOauthClient(userID int64, domain string) (client *models.Client, err error) {
	clientId := uuid.New().String()
	clientSecret := uuid.New().String()

	client = &models.Client{
		ID:     clientId,
		Secret: clientSecret,
		Domain: domain,
		UserID: IntToString(userID),
	}

	err = ClientStore.Create(client)
	if err != nil {
		return nil, err
	}

	return client, nil
}

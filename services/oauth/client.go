package oauth

import (
	. "google-scraper/helpers"

	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/google/uuid"
)


type ClientGenerator struct {
	Domain string
}

func (service *ClientGenerator) Generate(userID int64) (id string, err error) {
	clientId := uuid.New().String()
	clientSecret := uuid.New().String()

	client := &models.Client{
		ID:     clientId,
		Secret: clientSecret,
		Domain: service.Domain,
		UserID: IntToString(userID),
	}

	err = ClientStore.Create(client)
	if err != nil {
		return "", err
	}

	return client.GetID(), nil
}

package tests

import (
	"context"
	"time"

	"google-scraper/services/oauth"
	. "google-scraper/helpers"

	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo"
)

func FabricateOAuthToken(client *models.Client, userID int64) oauth2.TokenInfo {
	info := &models.Token{
		ClientID:        client.GetID(),
		UserID:         IntToString(userID),
		Access:          uuid.New().String(),
		AccessCreateAt:  time.Now().Local(),
		AccessExpiresIn: time.Second * 3600,
	}
	err := oauth.TokenStore.Create(context.Background(), info)
	if err != nil {
		Fail("Creating token failed:" + err.Error())
	}

	token, err := oauth.TokenStore.GetByAccess(context.Background(), info.GetAccess())
	if err != nil {
		Fail("Getting token info failed:" + err.Error())
	}

	return token
}

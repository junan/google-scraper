package initializers

import (
	"context"
	"time"

	"google-scraper/services/oauth"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/jackc/pgx/v4"
	pg "github.com/vgarvardt/go-oauth2-pg/v4"
	"github.com/vgarvardt/go-pg-adapter/pgx4adapter"
)

func init() {
	dbUrl, err := web.AppConfig.String("dbUrl")
	if err != nil {
		logs.Critical("Postgres database source is not found: ", err)
	}

	pgxConn, err := pgx.Connect(context.TODO(), dbUrl)
	if err != nil {
		logs.Critical("Connecting to Postgres failed: ", err)
	}

	manager := manage.NewDefaultManager()

	// use PostgreSQL token store with pgx.Connection adapter
	adapter := pgx4adapter.NewConn(pgxConn)
	tokenStore, err := pg.NewTokenStore(adapter, pg.WithTokenStoreGCInterval(time.Minute))
	if err != nil {
		logs.Critical("Creating token store failed: ", err)
	}
	defer tokenStore.Close()

	clientStore, err := pg.NewClientStore(adapter)
	if err != nil {
		logs.Critical("Creating client store failed: ", err)
	}

	manager.MapTokenStorage(tokenStore)
	manager.MapClientStorage(clientStore)

	oauthServer := server.NewDefaultServer(manager)
	oauthServer.SetAllowGetAccessRequest(true)
	oauthServer.SetClientInfoHandler(server.ClientFormHandler)
	oauthServer.SetPasswordAuthorizationHandler(oauth.PasswordAuthorizationHandler)
	oauthServer.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		logs.Critical("Setting internal error handler failed: ", err)
		return
	})
	oauthServer.SetResponseErrorHandler(func(re *errors.Response) {
		logs.Critical("Setting error response handler failed: ", err)
	})

	oauth.ClientStore = clientStore
	oauth.OauthServer = oauthServer
}

package oauth

import (
	"github.com/go-oauth2/oauth2/v4/server"
	pg "github.com/vgarvardt/go-oauth2-pg/v4"
)

var ClientStore *pg.ClientStore
var OauthServer *server.Server

func PasswordAuthorizationHandler(email string, password string) (string, error) {
	//sessionForm := forms.SessionForm{
	//	Email:    email,
	//	Password: password,
	//}
	//user, err := sessionForm.Authenticate()
	//if err != nil {
	//	return "", errors.ErrInvalidClient
	//}

	return "5", nil
}


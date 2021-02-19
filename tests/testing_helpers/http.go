package testing_helpers

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"

	"google-scraper/controllers"
	"google-scraper/models"

	"github.com/beego/beego/v2/server/web"
	"github.com/onsi/ginkgo"
)

func MakeRequest(method string, url string, body io.Reader) *http.Response {
	request, err := http.NewRequest(method, url, body)

	if err != nil {
		ginkgo.Fail("Create request failed: " + err.Error())
	}

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	response := httptest.NewRecorder()
	web.BeeApp.Handlers.ServeHTTP(response, request)

	return response.Result()
}

func MakeAuthenticatedRequest(method string, url string, body io.Reader, user *models.User) *http.Response {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		ginkgo.Fail("Failed to create request: " + err.Error())
	}

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	responseRecorder := httptest.NewRecorder()
	store, err := web.GlobalSessions.SessionStart(responseRecorder, request)
	if err != nil {
		ginkgo.Fail("Failed to start session" + err.Error())
	}

	err = store.Set(context.Background(), controllers.CurrentUserSession, user.Id)

	if err != nil {
		ginkgo.Fail("Failed to set current user" + err.Error())
	}

	web.BeeApp.Handlers.ServeHTTP(responseRecorder, request)

	return responseRecorder.Result()
}

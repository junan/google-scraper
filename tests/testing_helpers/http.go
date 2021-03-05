package testing_helpers

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"google-scraper/controllers"
	"google-scraper/models"
	. "google-scraper/services/crawler"

	"github.com/beego/beego/v2/server/web"
	"github.com/jarcoal/httpmock"
	. "github.com/onsi/ginkgo"
)

func MakeRequest(method string, url string, body io.Reader) *http.Response {
	request, err := http.NewRequest(method, url, body)

	if err != nil {
		Fail("Create request failed: " + err.Error())
	}

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	response := httptest.NewRecorder()
	web.BeeApp.Handlers.ServeHTTP(response, request)

	return response.Result()
}

func MakeAuthenticatedRequest(method string, url string, body io.Reader, contentType string,user *models.User) *http.Response {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		Fail("Failed to create request: " + err.Error())
	}

	request.Header.Add("Content-Type", contentType)

	responseRecorder := httptest.NewRecorder()
	store, err := web.GlobalSessions.SessionStart(responseRecorder, request)
	if err != nil {
		Fail("Failed to start session" + err.Error())
	}

	err = store.Set(context.Background(), controllers.CurrentUserSession, user.Id)

	if err != nil {
		Fail("Failed to set current user" + err.Error())
	}

	web.BeeApp.Handlers.ServeHTTP(responseRecorder, request)

	return responseRecorder.Result()
}

func MockCrawling(searchString string, filePath string) {
	searchUrl, err := BuildSearchUrl(searchString)
	if err != nil {
		Fail("Building search url failed: " + err.Error())
	}

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		Fail("Reading file failed: " + err.Error())
	}

	httpmock.RegisterResponder("GET", searchUrl,
		httpmock.NewStringResponder(200, string(content)))
}

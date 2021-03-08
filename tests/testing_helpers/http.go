package testing_helpers

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"

	"github.com/beego/beego/v2/server/web"
	"github.com/jarcoal/httpmock"
	. "github.com/onsi/ginkgo"

	"google-scraper/controllers"
	"google-scraper/models"
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

func MakeAuthenticatedRequest(method string, url string, body io.Reader, user *models.User) *http.Response {
	var contentType string
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		Fail("Failed to create request: " + err.Error())
	}

	if url == "/search" {
		contentType = "multipart/form-data; boundary=multipart-boundary"
	} else {
		contentType = "application/x-www-form-urlencoded"
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

func MockCrawling(mockResponseFilePath string) {
	content, err := ioutil.ReadFile(mockResponseFilePath)
	if err != nil {
		Fail("Reading file failed: " + err.Error())
	}

	httpmock.RegisterResponder("GET", `=~^https://www.google.com/search+\z`,
		httpmock.NewStringResponder(200, string(content)))
}

func CreateMultipartFormData(filePath string) (http.Header, *bytes.Buffer) {
	file, err := os.Open(filePath)
	if err != nil {
		Fail("Opening file failed: " + err.Error())
	}

	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	err = writer.SetBoundary("multipart-boundary")
	if err != nil {
		Fail("Setting multipart-boundary failed: " + err.Error())
	}

	writer.Close()
	part, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		Fail("Creating form file failed: " + err.Error())
	}

	_, err = io.Copy(part, file)
	if err != nil {
		Fail("Copying file failed: " + err.Error())
	}

	err = writer.Close()
	if err != nil {
		Fail("Closing writer failed: " + err.Error())
	}

	headers := http.Header{}
	headers.Set("Content-Type", writer.FormDataContentType())

	return headers, body
}


func GetFormFileData(filePath string) (multipart.File, *multipart.FileHeader, error) {
	headers, body := CreateMultipartFormData(filePath)

	req, err := http.NewRequest("POST", "", body)
	if err != nil {
		Fail("Creating request failed: " + err.Error())
	}
	req.Header = headers

	file, header, err := req.FormFile("file")
	if err != nil {
		Fail("Getting FormFile data failed: " + err.Error())
	}

	return file, header, err
}

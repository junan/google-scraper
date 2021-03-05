package testing_helpers

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/beego/beego/v2/server/web"
	. "github.com/onsi/ginkgo"
)

func GetSession(cookies []*http.Cookie, key string) interface{} {
	context := context.Background()

	for _, cookie := range cookies {
		if cookie.Name == web.BConfig.WebConfig.Session.SessionName {
			store, err := web.GlobalSessions.GetSessionStore(cookie.Value)
			if err != nil {
				Fail("Getting store failed: " + err.Error())
			}

			return store.Get(context, key)
		}
	}
	return nil
}

func CreateMultipartFormData(filename string) (*bytes.Buffer, string) {
	rootPath, err := os.Getwd()
	if err != nil {
		Fail("Getting rootPath failed: " + err.Error())
	}

	path := rootPath + "/tests/fixtures/controller/search/" + filename
	file, err := os.Open(path)
	if err != nil {
		Fail("Opening file failed: " + err.Error())
	}

	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", filepath.Base(path))

	_, err = io.Copy(part, file)
	if err != nil {
		Fail("Copying file failed: " + err.Error())
	}

	err = writer.Close()
	if err != nil {
		Fail("Closing writer failed: " + err.Error())
	}

	return body, writer.FormDataContentType()
}

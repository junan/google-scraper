package tests

import (
	"bytes"
	"context"
	"mime/multipart"
	"net/http"

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

func CreateEmptyMultipartBody() *bytes.Buffer {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	err := writer.SetBoundary("multipart-boundary")
	if err != nil {
		Fail("Setting multipart-boundary failed: " + err.Error())
	}

	writer.Close()

	return body
}

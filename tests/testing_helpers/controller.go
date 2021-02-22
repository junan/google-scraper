package testing_helpers

import (
	"net/http"
	"context"

	"github.com/beego/beego/v2/server/web"
	"github.com/onsi/ginkgo"
)

func GetSession(cookies []*http.Cookie, key string) interface{} {
	context := context.Background()

	for _, cookie := range cookies {
		if cookie.Name == web.BConfig.WebConfig.Session.SessionName {
			store, err := web.GlobalSessions.GetSessionStore(cookie.Value)
			if err != nil {
				ginkgo.Fail("Getting store failed: " + err.Error())
			}

			return store.Get(context, key)
		}
	}
	return nil
}

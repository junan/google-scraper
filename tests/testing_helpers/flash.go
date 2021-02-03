package test_helpers

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/beego/beego/v2/server/web"
)

func GetFlash(cookies []*http.Cookie) *web.FlashData {
	for _, cookie := range cookies {
		if cookie.Name == "BEEGO_FLASH" {
			return readFromCookie(cookie)
		}
	}

	return nil
}

func readFromCookie(cookie *http.Cookie) *web.FlashData {
	flash := web.NewFlash()
	if cookie != nil {
		v, _ := url.QueryUnescape(cookie.Value)
		vals := strings.Split(v, "\x00")
		for _, v := range vals {
			if len(v) > 0 {
				kv := strings.Split(v, "\x23"+web.BConfig.WebConfig.FlashSeparator+"\x23")
				if len(kv) == 2 {
					flash.Data[kv[0]] = kv[1]
				}
			}
		}
	}
	return flash
}

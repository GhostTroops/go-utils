package go_utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func Fanyi4Youdao(sT string) {
	x := "webdict"
	v := sT
	t := v + x
	time1 := len(v+x) % 10
	s := fmt.Sprintf("%x", md5.Sum([]byte(t)))
	s = "web" + v + fmt.Sprintf("%d", time1) + "Mk6hqtUp33DGGtoS63tTJbMUYjRrG1Lu" + s
	s = fmt.Sprintf("%x", md5.Sum([]byte(s)))
	data := []byte("q=" + url.QueryEscape(v) + "&le=en&t=2&client=web&sign=" + s + "&keyfrom=" + x)
	DoUrlCbk4byte("https://dict.youdao.com/jsonapi_s?doctype=json&jsonversion=4", data, map[string]string{
		"Accept":             "application/json, text/plain, */*",
		"Accept-Language":    "zh,en;q=0.9,zh-CN;q=0.8",
		"Connection":         "keep-alive",
		"Content-Type":       "application/x-www-form-urlencoded",
		"DNT":                "1",
		"Origin":             "https://www.youdao.com",
		"Referer":            "https://www.youdao.com/",
		"Sec-Fetch-Dest":     "empty",
		"Sec-Fetch-Mode":     "cors",
		"Sec-Fetch-Site":     "same-site",
		"User-Agent":         "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36",
		"sec-ch-ua":          `"Chromium";v="118", "Google Chrome";v="118", "Not=A?Brand";v="99"`,
		"sec-ch-ua-mobile":   "?0",
		"sec-ch-ua-platform": "macOS",
	}, func(resp *http.Response, szUrl string) {
		if data1, err := io.ReadAll(resp.Body); nil == err {
			var m = map[string]interface{}{}
			if Json.Unmarshal(data1, &m) == nil {
				s = GetJQ2Str(m, ".fanyi.tran")
				fmt.Println(s)
			}
		}
	})
}

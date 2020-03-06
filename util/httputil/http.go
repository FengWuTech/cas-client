package httputil

import (
	"cas-client/util/logutil"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func BuildURL(baseURL string, param map[string]interface{}) string {
	params := url.Values{}
	for k, v := range param {
		switch v.(type) {
		case string:
			params.Set(k, v.(string))
		case int:
			params.Set(k, strconv.Itoa(v.(int)))
		default:
			panic("param存在不支持的参数类型")
		}
	}

	parseURL, _ := url.Parse(baseURL)
	if len(parseURL.Query()) == 0 {
		return baseURL + "?" + params.Encode()
	} else {
		return baseURL + "&" + params.Encode()
	}
}

func Get(url string) (error, []byte) {
	var start = time.Now().UnixNano()
	resp, err := http.Get(url)
	if err != nil {
		logutil.WarnF("url[%v] error[%v]", url, err)
		return err, nil
	}
	resBody, _ := ioutil.ReadAll(resp.Body)
	var end = time.Now().UnixNano()
	spend := (end - start) / int64(time.Millisecond)
	logutil.InfoF("%dms url[%v] res[%v]", spend, url, string(resBody))
	return nil, resBody
}

func PostRawJson(url string, body string) (error, []byte) {
	var start = time.Now().UnixNano()
	resp, err := http.Post(url, "application/json", strings.NewReader(body))
	if err != nil {
		logutil.WarnF("url[%v] body[%v] error[%v]", url, body, err)
		return err, nil
	}
	resBody, err := ioutil.ReadAll(resp.Body)
	var end = time.Now().UnixNano()
	spend := (end - start) / int64(time.Millisecond)
	logutil.InfoF("%dms url[%v] body[%v] res[%v] error[%v]", spend, url, body, string(resBody), err)
	return nil, resBody
}

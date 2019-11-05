package helper

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kataras/iris"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)


// 获取iris的请求头
func RequestHeader(ctx iris.Context) string {
	var requestHeader string
	for k, v := range ctx.Request().Header {
		requestHeader += k + "=" + v[0] + ";"
	}
	return requestHeader
}

// 获取iris的请求体
func RequestBody(ctx iris.Context) string {
	var requestBody string

	// 如果不是json格式的，则不解析
	if ctx.GetHeader("Content-Type") != "application/json" {
		return requestBody
	}

	data, err := ioutil.ReadAll(ctx.Request().Body)
	if err == nil {
		requestBody = string(data)
		ctx.Request().Body = ioutil.NopCloser(bytes.NewBuffer(data))
	}

	return requestBody
}

// 获取iris的get参数
func RequestQueries(ctx iris.Context) string {
	var requestQuery string
	for k, v := range ctx.URLParams() {
		requestQuery += k + "=" + v + "&"
	}
	requestQuery = strings.Trim(requestQuery, "&")

	return requestQuery
}

// 协程共享的
var httpClient = &http.Client{
	Transport: &http.Transport{
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		DisableCompression:  true,
		MaxIdleConnsPerHost: 50,
	},
	Timeout: time.Duration(30) * time.Second,
}

func Post(url string, header map[string]string, data map[string]interface{}) ([]byte, error) {
	var err error
	var req *http.Request
	var resp *http.Response

	if b, err := json.Marshal(data); err != nil {
		return nil, err
	} else {
		if req, err = http.NewRequest("POST", url, bytes.NewReader(b)); err != nil {
			return nil, err
		}
	}

	if len(header) > 0 {
		for key, value := range header {
			req.Header.Add(key, value)
		}
	}

	if resp, err = httpClient.Do(req); err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("ErrCode:%d, with %s => %#v", resp.StatusCode, url, data))
	}

	if body, err := ioutil.ReadAll(resp.Body); err != nil {
		return nil, err
	} else {
		return body, nil
	}
}

func Get(url string, header map[string]string, data map[string]interface{}) ([]byte, error) {
	var err error
	var req *http.Request
	var resp *http.Response

	if req, err = http.NewRequest("GET", url, nil); err != nil {
		return nil, err
	}

	if len(data) > 0 {
		query := req.URL.Query()
		for k, v := range data {
			query.Add(k, fmt.Sprint(v))
		}

		req.URL.RawQuery = query.Encode()
	}

	if len(header) > 0 {
		for key, value := range header {
			req.Header.Add(key, value)
		}
	}

	if resp, err = httpClient.Do(req); err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("ErrCode:%d, with %s => %#v", resp.StatusCode, url, data))
	}

	if body, err := ioutil.ReadAll(resp.Body); err != nil {
		return nil, err
	} else {
		return body, nil
	}
}

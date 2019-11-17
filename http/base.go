package http

import (
	"demo/utility/helper"
	"demo/utility/log"
	"fmt"
	"go.uber.org/zap"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type Response interface {
	GetCode() int
	SetData(interface{})
	String() string
	Translate(string)
}

type response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg,omitempty"`
	Tip  string      `json:"tip,omitempty"` // 用来获取调式信息
	Data interface{} `json:"data,omitempty"`
}

func (resp *response) String() string {
	return fmt.Sprintf("code: %d, msg: %s, tip: %s", resp.Code, resp.Msg, resp.Tip)
}

func (resp *response) GetCode() int {
	return resp.Code
}

func (resp *response) SetData(data interface{}) {
	resp.Data = data
}

func (resp *response) Translate(lang string) {
	// Todo 翻译
	resp.Msg = fmt.Sprint(resp.Code)
}

func newResponse(code int, msg string, options ...interface{}) Response {
	resp := &response{Code: code, Msg: msg}

	if code != RESP_SUCCESS {
		if _, file, line, ok := runtime.Caller(2); ok {
			fileAndLine := strings.ReplaceAll(filepath.Base(file), ".go", strconv.Itoa(line))
			resp.Tip = helper.MD5(fmt.Sprint(time.Now().UnixNano())) + " -- " + fileAndLine
		}
	}

	if len(options) > 0 {
		var errs, tips string

		for _, option := range options {
			if err, ok := option.(error); ok {
				errs += err.Error() + ";"
			} else {
				tips += fmt.Sprint(option)
			}
		}

		if tips != "" {
			resp.Tip += " -- " + tips
		}

		if errs != "" {
			log.Error(resp.String(), zap.String("error", errs))
		}
	}

	return resp
}

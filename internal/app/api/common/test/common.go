package test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/urunimi/gorest/core"
)

// ControllerSuite struct definition
type ControllerSuite struct {
}

func (ts *ControllerSuite) BuildParams(params map[string]interface{}) *url.Values {
	vals := url.Values{
		"deviceModel": []string{"SM-G930T"},
		"buildId":     []string{"NRD90M.G930TUVS4BQJ2"},
		"sdkVersion":  []string{"24"},
		"versionCode": []string{"4088"},
		"locale":      []string{"en_US"},
	}
	if params != nil {
		for k, v := range params {
			vals[k] = []string{fmt.Sprintf("%v", v)}
		}
	}
	return &vals
}

func (ts *ControllerSuite) BuildContextAndRecorder(httpRequest *http.Request) (ctx core.Context, rec *httptest.ResponseRecorder) {
	engine := core.NewEngine()
	rec = httptest.NewRecorder()
	ctx = engine.NewContext(httpRequest, rec)
	return
}

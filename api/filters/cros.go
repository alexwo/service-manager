package filters

import (
	"github.com/Peripli/service-manager/pkg/env"
	"github.com/Peripli/service-manager/pkg/web"
	"net/http"
)

// ResponseHeaderStripperFilter is a web.Filter used to strip headers from all OSB calls
// on their way back to a platform
type CROSFilter struct {
	Environment env.Environment
	web.Filter
	Headers []string
}

// Name implements web.Named and returns the Filter name
func (f *CROSFilter) Name() string {
	return "CROSFilter"
}

func resWrap(res *web.Response, err error, allowedUrl string) (*web.Response, error) {
	res.Header.Add("Access-Control-Allow-Origin", allowedUrl)
	return res, err
}

// Run implements web.Filter and represents the Response Header Stripper middleware function that
// strips blacklisted headers
func (f *CROSFilter) Run(request *web.Request, next web.Handler) (*web.Response, error) {
	var allowedUrl string
	reqHost := request.Header.Get("Origin")
	allowedHost := f.Environment.Get("cross.allowed_host")
	if allowedHost != nil {
		allowedUrl = allowedHost.(string)
	}

	var webRes web.Response
	webRes.StatusCode = 405
	webRes.Header = http.Header{}
	if request.Method == "OPTIONS" {
		if allowedHost == reqHost {
			webRes.Header.Add("Access-Control-Allow-Origin", allowedUrl)
			webRes.Header.Add("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			webRes.Header.Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			webRes.StatusCode = 200
			return &webRes, nil
		} else {
			return &webRes, nil
		}
	} else if allowedHost == reqHost {
		res, err := next.Handle(request)
		if err == nil {
			return resWrap(res, err, allowedUrl)
		} else {
			return res, err
		}
	}

	return next.Handle(request)
}

// FilterMatchers implements the web.Filter interface and returns the conditions
// on which the filter should be executed
func (f *CROSFilter) FilterMatchers() []web.FilterMatcher {
	return []web.FilterMatcher{
		{
			Matchers: []web.Matcher{
				web.Path("*/**"),
				web.Methods(http.MethodOptions, http.MethodGet, http.MethodPost),
			},
		},
	}
}

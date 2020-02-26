package filters

import (
	"github.com/Peripli/service-manager/pkg/web"
	"net/http"
	"strings"
	"github.com/Peripli/service-manager/pkg/env"
	"fmt"
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

// Run implements web.Filter and represents the Response Header Stripper middleware function that
// strips blacklisted headers
func (f *CROSFilter) Run(request *web.Request, next web.Handler) (*web.Response, error) {

	reqHost := request.Header.Get("host")
	var webRes web.Response;
	webRes.StatusCode = 405
	if request.Method == "OPTIONS" {

		allowedHosts := f.Environment.Get("cross.allowed_hosts")
		fmt.Println("Came from hist %s", reqHost)
		fmt.Println("allowedHost: %s", allowedHosts)
		if allowedHosts != nil {
			allowedHosts := allowedHosts.(string)
			allowedHostsArray := strings.Split(allowedHosts, ",")

			for _, allowedHost := range allowedHostsArray {
				if strings.ContainsAny(allowedHost, reqHost) {
					webRes.Header = http.Header{}
					webRes.Header.Add("Access-Control-Allow-Origin", "*")
					webRes.Header.Add("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
					webRes.Header.Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
					webRes.StatusCode = 200
					return &webRes, nil
				}
			}
		}
		return &webRes, nil
	}

	return &webRes, nil
}

// FilterMatchers implements the web.Filter interface and returns the conditions
// on which the filter should be executed
func (f *CROSFilter) FilterMatchers() []web.FilterMatcher {
	return []web.FilterMatcher{
		{
			Matchers: []web.Matcher{
				web.Path("*/**"),
				web.Methods(http.MethodOptions),
			},
		},
	}
}

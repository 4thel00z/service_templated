package debug

import (
	"service_templated/pkg/libservice_template"
)

type Debug struct{}

var (
	Module = Debug{}
)

func (d Debug) Version() string {
	return "v1"
}

func (d Debug) Namespace() string {
	return "debug"
}

func (d Debug) Routes() map[string]libservice_template.Route {
	// Add route definitions here
	return map[string]libservice_template.Route{
		"routes": {
			Path:        "routes",
			Method:      "GET",
			CurlExample: "curl http://<addr>/<version>/<namespace>/<path>",
			Service:     GetRoutesHandler,
		},
	}
}

func (d Debug) LongPath(route libservice_template.Route) string {
	return libservice_template.DefaultLongPath(d, route)
}

package debug

import (
	libservice "github.com/4thel00z/libservice/v1"
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

func (d Debug) Routes() map[string]libservice.Route {
	// Add route definitions here
	return map[string]libservice.Route{
		"routes": {
			Path:        "routes",
			Method:      "GET",
			CurlExample: "curl http://<addr>/<version>/<namespace>/<path>",
			Service:     GetRoutesHandler,
		},
	}
}

func (d Debug) LongPath(route libservice.Route) string {
	return libservice.DefaultLongPath(d, route)
}

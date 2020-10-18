package debug

import "service_templated/pkg/libservice_template"

type Debug struct{}

var (
	Module = Debug{}
)

func (Y Debug) Version() string {
	return "v1"
}

func (Y Debug) Namespace() string {
	return "debug"
}

func (Y Debug) Routes() map[string]libservice_template.Route {
	// Add route definitions here
	return map[string]libservice_template.Route{
		"routes": {
			Path:        "routes",
			Method:      "GET",
			CurlExample: "curl http://<addr>/<version>/<namespace>/routes",
			Service:     GetRoutesHandler,
		},
	}
}

func (Y Debug) LongPath(route libservice_template.Route) string {
	return libservice_template.DefaultLongPath(Y, route)
}

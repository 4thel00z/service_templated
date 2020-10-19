package debug

import (
	"service_templated/pkg/libservice_template"
	"service_templated/pkg/libservice_template/jwt"
)

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
		"private": {
			Path:        "private",
			Method:      "GET",
			CurlExample: "curl http://<addr>/<version>/<namespace>/private",
			Service:     GetPrivateMessageHandler,
			TokenValidator: jwt.New(
				jwt.WithDebug(),
			).Middleware,
		},
	}
}

func (Y Debug) LongPath(route libservice_template.Route) string {
	return libservice_template.DefaultLongPath(Y, route)
}

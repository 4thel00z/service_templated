package examples

import (
	"service_templated/pkg/libservice_template"
	"service_templated/pkg/libservice_template/jwt"
)

type Example struct{}

var (
	Module = Example{}
)

func (e Example) Version() string {
	return "v1"
}

func (e Example) Namespace() string {
	return "examples"
}

func (e Example) Routes() map[string]libservice_template.Route {
	// Add route definitions here
	return map[string]libservice_template.Route{
		"private": {
			Path:        "private",
			Method:      "GET",
			CurlExample: "curl http://<addr>/<version>/<namespace>/<path>",
			Service:     GetPrivateMessageHandler,
			TokenValidator: jwt.New(
				jwt.WithDebug(),
			).Middleware,
		},
		"multipart": {
			Path:        "multipart",
			Method:      "POST",
			CurlExample: "curl -F pdf_file=@examples/some.pdf http://<addr>/<version>/<namespace>/<path>",
			Service:     PostMultiPartMessageHandler,
			// Validate with the PostMultiPartMessage tags and accept files up to 10mb
			MultiPartValidator: libservice_template.GenerateMultipartValidator(PostMultiPartMessage{}, 10<<20),
		},
	}
}

func (e Example) LongPath(route libservice_template.Route) string {
	return libservice_template.DefaultLongPath(e, route)
}

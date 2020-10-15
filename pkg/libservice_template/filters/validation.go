package filters

import (
	"fmt"
	"github.com/monzo/typhon"
	"service_templated/pkg/libservice_template"
)

func Validation(app libservice_template.App) typhon.Filter {
	return func(req typhon.Request, svc typhon.Service) typhon.Response {
		pattern := app.Router.Pattern(req)
		for _, route := range app.Routes() {
			if route.LongPath != pattern {
				continue
			}

			if route.Validator == nil {
				return svc(req)
			}

			err := (*route.Validator)(req)

			if err != nil {
				msg := err.Error()
				return req.Response(libservice_template.GenericResponse{
					Message: fmt.Sprintf("[%s] %s validation error", route.LongPath, route.Method),
					Error:   &msg,
				})
			}
		}

		return svc(req)
	}
}

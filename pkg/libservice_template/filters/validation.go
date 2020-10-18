package filters

import (
	"context"
	"fmt"
	"github.com/monzo/typhon"
	"service_templated/pkg/libservice_template"
)

const (
	ValidationResult = "validation_result"
)

func Validation(app libservice_template.App) typhon.Filter {
	return func(req typhon.Request, svc typhon.Service) typhon.Response {
		pattern := app.Router.Pattern(req)
		routes := app.Routes()
		route, ok := routes[pattern]
		if !ok {
			return svc(req)
		}

		if route.Validator == nil {
			return svc(req)
		}

		val, err := (*route.Validator)(req)

		if err != nil {
			msg := err.Error()
			return req.Response(libservice_template.GenericResponse{
				Message: fmt.Sprintf("[%s] %s validation error", pattern, route.Method),
				Error:   &msg,
			})
		}

		req.Context = context.WithValue(req.Context, ValidationResult, val)
		return svc(req)

	}
}

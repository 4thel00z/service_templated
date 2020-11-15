package filters

import (
	"context"
	"fmt"
	"github.com/monzo/typhon"
	"service_templated/pkg/libservice_template"
)

const (
	MultipartValidationResult = "multipart_validation_result"
)

// TODO: add file and parameter support at the same time, probably by making the tag parsing more complex
// i.e. add other flags than "required" like "is_file", etc.		 
func MultipartValidation(app libservice_template.App) typhon.Filter {
	return func(req typhon.Request, svc typhon.Service) typhon.Response {
		pattern := app.Router.Pattern(req)
		routes := app.Routes()
		route, ok := routes[pattern]
		if !ok {
			return svc(req)
		}

		if route.MultiPartValidator == nil {
			return svc(req)
		}

		val, err := (*route.MultiPartValidator)(req)

		if err != nil {
			msg := err.Error()
			return req.Response(libservice_template.GenericResponse{
				Message: fmt.Sprintf("[%s] %s validation error", pattern, route.Method),
				Error:   &msg,
			})
		}

		req.Context = context.WithValue(req.Context, MultipartValidationResult, val)
		return svc(req)

	}
}

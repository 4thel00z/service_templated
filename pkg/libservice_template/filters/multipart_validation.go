package filters

import (
	"context"
	"fmt"
	"github.com/4thel00z/libhttp"

	libservice "github.com/4thel00z/libservice/v1"
)

const (
	MultipartValidationResult = "multipart_validation_result"
)

// TODO: add file and parameter support at the same time, probably by making the tag parsing more complex
// i.e. add other flags than "required" like "is_file", etc.		 
func MultipartValidation(app libservice.App) libhttp.Filter {
	return func(req libhttp.Request, svc libhttp.Service) libhttp.Response {
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
			return req.Response(libservice.GenericResponse{
				Message: fmt.Sprintf("[%s] %s validation error", pattern, route.Method),
				Error:   &msg,
			})
		}

		req.Context = context.WithValue(req.Context, MultipartValidationResult, val)
		return svc(req)

	}
}

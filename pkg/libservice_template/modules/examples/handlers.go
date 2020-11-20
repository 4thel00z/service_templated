package examples

import (
	"fmt"
	"github.com/4thel00z/libhttp"
	"github.com/dgrijalva/jwt-go"

	libservice "github.com/4thel00z/libservice/v1"
	"service_templated/pkg/libservice_template/filters"
	libjwt "service_templated/pkg/libservice_template/jwt"
)

func PostMultiPartMessageHandler(app libservice.App) libhttp.Service {
	return func(req libhttp.Request) libhttp.Response {
		multi := req.Value(filters.MultipartValidationResult).(*PostMultiPartMessage)
		response := req.Response(&libservice.GenericResponse{
			Message: string(multi.File),
		})
		response.StatusCode = 200
		return response
	}
}

func GetPrivateMessageHandler(app libservice.App) libhttp.Service {
	return func(req libhttp.Request) libhttp.Response {
		token := req.Value(libjwt.DefaultUserProperty).(*jwt.Token)
		response := req.Response(&libservice.GenericResponse{
			Message: fmt.Sprintf("This is my token: %s!", token.Raw),
		})
		response.StatusCode = 200
		return response
	}
}

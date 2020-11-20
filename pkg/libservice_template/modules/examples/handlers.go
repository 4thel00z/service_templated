package examples

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/monzo/typhon"
	"service_templated/pkg/libservice"
	"service_templated/pkg/libservice_template/filters"
	libjwt "service_templated/pkg/libservice_template/jwt"
)

func PostMultiPartMessageHandler(app libservice.App) typhon.Service {
	return func(req typhon.Request) typhon.Response {
		multi := req.Value(filters.MultipartValidationResult).(*PostMultiPartMessage)
		response := req.Response(&libservice.GenericResponse{
			Message: string(multi.File),
		})
		response.StatusCode = 200
		return response
	}
}

func GetPrivateMessageHandler(app libservice.App) typhon.Service {
	return func(req typhon.Request) typhon.Response {
		token := req.Value(libjwt.DefaultUserProperty).(*jwt.Token)
		response := req.Response(&libservice.GenericResponse{
			Message: fmt.Sprintf("This is my token: %s!", token.Raw),
		})
		response.StatusCode = 200
		return response
	}
}

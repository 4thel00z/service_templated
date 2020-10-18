package debug

import (
	"github.com/monzo/typhon"
	"service_templated/pkg/libservice_template"
)

func GetRoutesHandler(app libservice_template.App) typhon.Service {
	return func(req typhon.Request) typhon.Response {

		response := req.Response(&GetRoutesResponse{
			Routes: app.Routes(),
		})

		response.StatusCode = 200
		return response
	}
}

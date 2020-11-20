package debug

import (
	"github.com/4thel00z/libhttp"
	libservice "github.com/4thel00z/libservice/v1"
)

func GetRoutesHandler(app libservice.App) libhttp.Service {
	return func(req libhttp.Request) libhttp.Response {

		response := req.Response(&GetRoutesResponse{
			Routes: app.Routes(),
		})

		response.StatusCode = 200
		return response
	}
}

package debug

import (
	libservice "github.com/4thel00z/libservice/v1"
)

type GetRoutesResponse struct {
	Routes map[string]libservice.Route `json:"routes"`
	Error  *string                     `json:"error,omitempty"`
}

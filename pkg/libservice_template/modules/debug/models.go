package debug

import (
	"service_templated/pkg/libservice"
)

type GetRoutesResponse struct {
	Routes map[string]libservice.Route `json:"routes"`
	Error  *string                     `json:"error,omitempty"`
}

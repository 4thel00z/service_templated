package debug

import "service_templated/pkg/libservice_template"

type GetRoutesResponse struct {
	Routes map[string]libservice_template.Route `json:"routes"`
	Error  *string                              `json:"error,omitempty"`
}

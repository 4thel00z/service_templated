package debug

import "service_templated/pkg/libservice_template"

type GenericResponse struct {
	Message interface{} `json:"message"`
	Error   *string     `json:"error,omitempty"`
}

type GetRoutesResponse struct {
	Routes []libservice_template.Route `json:"routes"`
	Error  *string         `json:"error,omitempty"`
}

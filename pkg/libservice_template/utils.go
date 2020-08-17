package libservice_template

import (
	"fmt"
	"github.com/monzo/typhon"
)

func Default404Handler(app App) typhon.Service {
	return func(req typhon.Request) typhon.Response {
		// TODO: Change this body to a default 404 page
		response := req.Response(nil)
		response.StatusCode = 404
		return response
	}
}

func VerifyRequest(r *typhon.Request) error {
	if _, ok := r.MultipartForm.File["audio"]; !ok || len(r.MultipartForm.File["audio"]) < 1 {
		return fmt.Errorf("the field audio is absent")
	}
	return nil
}

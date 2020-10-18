package libservice_template

import (
	"encoding/json"
	"github.com/monzo/typhon"
	"io/ioutil"
	"os"
	"strings"
)

type GenericResponse struct {
	Message interface{} `json:"message"`
	Error   *string     `json:"error,omitempty"`
}

func ParseConfig(path string) (config Config, err error) {
	file, err := os.Open(path)
	if err != nil {
		return Config{}, err
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return Config{}, err
	}
	err = json.Unmarshal(content, &config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}

type Config struct {
	//TODO: add more fields here if you want to make the app more configurable
}

type Service func(app App) typhon.Service
type Validator func(request typhon.Request) (interface{}, error)
type TokenValidator func(request typhon.Request) (interface{}, error)

type Route struct {
	Path           string          `json:"-"`
	Method         string          `json:"method"`
	CurlExample    string          `json:"curl_example"`
	Validator      *Validator      `json:"-"`
	TokenValidator *TokenValidator `json:"-"`
	Service        Service         `json:"-"`
}

func NewRoute(path, method, curlExample string, validator *Validator, tokenValidator *TokenValidator, svc Service) Route {
	if svc == nil {
		svc = Default404Handler
	}

	return Route{
		Path:           path,
		Method:         method,
		CurlExample:    curlExample,
		Validator:      validator,
		TokenValidator: tokenValidator,
	}
}

type Module interface {
	Version() string
	Namespace() string
	Routes() map[string]Route
	LongPath(route Route) string
}

func DefaultLongPath(module Module, route Route) string {
	return "/" + strings.Join([]string{module.Version(), module.Namespace(), route.Path}, "/")
}

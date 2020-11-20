package jwt

import (
	"github.com/4thel00z/libhttp"
	"github.com/dgrijalva/jwt-go"
)

type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}

type JSONWebKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

type CustomClaims struct {
	Scope string `json:"scope"`
	jwt.StandardClaims
}

type TokenExtractor func(r libhttp.Request) (string, error)
type ErrorHandler func(r libhttp.Request, errMsg string) libhttp.Response
type EmptyTokenHandler libhttp.Service
type ScopeChecker func(tokenString string) bool

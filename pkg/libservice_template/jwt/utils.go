package auth

import (
	"encoding/json"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"strings"
)

const (
	OAuthAudienceEnvKey = "OAUTH_AUDIENCE"
	OAuthIssuerEnvKey   = "OAUTH_ISSUER"
	OAuthJWKSUrlEnvKey  = "OAUTH_JWKS_URL"
)

func CheckOAuthScope(jwksUrl, scope string) func(tokenString string) bool {
	return func(tokenString string) bool {
		token, _ := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			cert, err := getPemCert(jwksUrl, token)
			if err != nil {
				return nil, err
			}
			result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
			return result, nil
		})

		claims, ok := token.Claims.(*CustomClaims)

		hasScope := false
		if ok && token.Valid {
			result := strings.Split(claims.Scope, " ")
			for i := range result {
				if result[i] == scope {
					hasScope = true
				}
			}
		}

		return hasScope
	}
}

func CheckOAuthScopeFromEnv(scope string) func(tokenString string) bool {
	return func(tokenString string) bool {
		token, _ := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			cert, err := getPemCert(os.ExpandEnv(OAuthAudienceEnvKey), token)
			if err != nil {
				return nil, err
			}
			result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
			return result, nil
		})

		claims, ok := token.Claims.(*CustomClaims)

		hasScope := false
		if ok && token.Valid {
			result := strings.Split(claims.Scope, " ")
			for i := range result {
				if result[i] == scope {
					hasScope = true
				}
			}
		}

		return hasScope
	}
}
func ValidationKeyGetterFromEnv() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		if token == nil {
			return nil, errors.New("(token *jwt.Token) is nil inside of the ValidationKeyGetter")
		}

		checkAud := token.Claims.(jwt.MapClaims).VerifyAudience(os.ExpandEnv(OAuthAudienceEnvKey), false)
		if !checkAud {
			return token, errors.New("invalid audience")
		}

		checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(os.ExpandEnv(OAuthIssuerEnvKey), false)
		if !checkIss {
			return token, errors.New("invalid issuer")
		}

		cert, err := getPemCert(os.ExpandEnv(OAuthJWKSUrlEnvKey), token)
		if err != nil {
			return nil, err
		}

		result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
		return result, nil
	}
}

func ValidationKeyGetterFromMetaData(aud, iss, jwksUrl string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		if token == nil {
			return nil, errors.New("(token *jwt.Token) is nil inside of the ValidationKeyGetter")
		}

		checkAud := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)
		if !checkAud {
			return token, errors.New("invalid audience")
		}

		checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
		if !checkIss {
			return token, errors.New("invalid issuer")
		}

		cert, err := getPemCert(jwksUrl, token)
		if err != nil {
			return nil, err
		}

		result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
		return result, nil
	}
}

func getPemCert(jwksUrl string, token *jwt.Token) (string, error) {
	cert := ""
	resp, err := http.Get(jwksUrl)

	if err != nil {
		return cert, err
	}
	defer resp.Body.Close()

	var jwks = Jwks{}
	err = json.NewDecoder(resp.Body).Decode(&jwks)

	if err != nil {
		return cert, err
	}

	for k, _ := range jwks.Keys {
		if token.Header["kid"] == jwks.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + jwks.Keys[k].X5c[0] + "\n-----END CERTIFICATE-----"
		}
	}

	if cert == "" {
		return cert, errors.New("unable to find appropriate key")
	}

	return cert, nil
}

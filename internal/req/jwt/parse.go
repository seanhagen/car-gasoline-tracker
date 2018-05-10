package jwt

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/oauth2/google"
	oauthsvc "google.golang.org/api/oauth2/v2"
)

// ParseGoogleSvcAcct parses a Google Service Account JSON file to create a JWT
func ParseGoogleSvcAcct(path string) (*jwt.Token, error) {
	if path == "" {
		return nil, fmt.Errorf("need a path to service account")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("file %#v does not exist, %v", path, err)
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	tkn, err := google.JWTAccessTokenSourceFromJSON(data, "biba-services.endpoints.sample.google.com")
	if err != nil {
		return nil, err
	}

	x, err := tkn.Token()
	if err != nil {
		return nil, err
	}

	conf, err := google.JWTConfigFromJSON(data, oauthsvc.UserinfoEmailScope)
	if err != nil {
		return nil, err
	}

	y, err := jwt.Parse(x.AccessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
		}

		key, err := jwt.ParseRSAPrivateKeyFromPEM(conf.PrivateKey)
		if err != nil {
			return nil, fmt.Errorf("Unable to parse private key: %v", err)
		}

		return key.Public(), nil
	})

	if err != nil {
		return nil, err
	}

	return y, nil
}

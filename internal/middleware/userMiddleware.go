package userMiddleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
)

func GetUserId(r *http.Request) (UserId int, err error) {

	if accessToken, err := getAccessTokenFromHeader(r); err != nil {
		return 0, err
	}
	token, err := jwt.ParseWithClaims(accessToken, &auth.Claims{}, func(token *jwt.Token) (interface{}, error) {

	})

	return 0, nil
}

func getAccessTokenFromHeader(r *http.Request) (bearer string, err error) {
	var authHeader string = r.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("no_authorization")
	}

	var headerParts []string = strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		return "", errors.New("no_authorization")
	}

	if headerParts[0] != "Bearer" {
		return "", errors.New("no_authorization")
	}

	return headerParts[1], nil
}

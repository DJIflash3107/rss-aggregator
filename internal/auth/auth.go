package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authentication information found")
	}
	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed authentication")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of authentication")
	}
	return vals[1], nil
}
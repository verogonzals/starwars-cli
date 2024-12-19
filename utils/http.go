package utils

import (
	"net/http"
	"strings"
)

func InitHttpRequest(url, jwt, method, body string, queries map[string]string) (*http.Request, error) {

	req, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	return req, nil

}

package api

import (
	"fmt"
	"net/http"
)

type Authorization struct {
	Token string
}

func (authorization Authorization) doAuthorize(request *http.Request) {
	request.Header.Set("Authorization", fmt.Sprintf("Token %s", authorization.Token))
}

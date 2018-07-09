package api

import (
	"gopkg.in/guregu/null.v3"
	"net/http"
	"testing"
)

var doAuthorizeCases = []struct {
	in           *Authorization
	needsRequest bool
	expected     *string
}{
	{
		in: &Authorization{
			Token: "token",
		},
		needsRequest: true,
		expected:     null.StringFrom("Token token").Ptr(),
	},
	{
		in: &Authorization{
			Token: "token",
		},
		needsRequest: false,
		expected:     nil,
	},
	{
		in:           nil,
		needsRequest: true,
		expected:     nil,
	},
	{
		in:           nil,
		needsRequest: false,
		expected:     nil,
	},
}

func TestAuthorization_doAuthorize(t *testing.T) {
	for i, c := range doAuthorizeCases {
		t.Logf("TestAuthorization_doAuthorize at %d", i)

		var req *http.Request = nil

		if c.needsRequest {
			req, _ = http.NewRequest(http.MethodGet, "xxxx", nil)
		}

		c.in.doAuthorize(req)

		if c.expected != nil {
			req.Header.Get("Authorization")
		}
	}
}

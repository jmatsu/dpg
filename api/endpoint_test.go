package api

import (
	"testing"
)

// https://docs.deploygate.com/reference#upload

func TestUploadAppEndpointToURL(t *testing.T) {
	cases := []struct {
		in       UploadAppEndpoint
		expected string
	}{
		{
			in: UploadAppEndpoint{
				BaseURL:      "x",
				AppOwnerName: "y",
			},
			expected: "x/api/users/y/apps",
		},
	}

	for i, c := range cases {
		t.Logf("TestUploadAppEndpointToURL at %d", i)

		if c.in.ToURL() != c.expected {
			t.Errorf("%s was expected but %s was found.", c.in.ToURL(), c.expected)
		}
	}
}

// https://docs.deploygate.com/reference#invite

package api

import (
	"testing"
)

// https://docs.deploygate.com/reference#upload

func TestAppUploadEndpoint_ToURL(t *testing.T) {
	cases := []struct {
		in       AppsEndpoint
		expected string
	}{
		{
			in: AppsEndpoint{
				BaseURL:      "x",
				AppOwnerName: "y",
			},
			expected: "x/api/users/y/apps",
		},
	}

	for i, c := range cases {
		t.Logf("TestAppUploadEndpoint_ToURL at %d", i)

		if c.in.ToURL() != c.expected {
			t.Errorf("%s was expected but %s was found.", c.expected, c.in.ToURL())
		}
	}
}

// https://docs.deploygate.com/reference#invite

func TestAppMemberEndpoint_ToURL(t *testing.T) {
	cases := []struct {
		in       AppMembersEndpoint
		expected string
	}{
		{
			in: AppMembersEndpoint{
				BaseURL:      "x",
				AppOwnerName: "y",
				AppPlatform:  "android",
				AppId:        "test package",
			},
			expected: "x/api/users/y/platforms/android/apps/test package/members",
		},
	}

	for i, c := range cases {
		t.Logf("TestAppMemberEndpoint_ToURL at %d", i)

		if c.in.ToURL() != c.expected {
			t.Errorf("%s was expected but %s was found.", c.expected, c.in.ToURL())
		}
	}
}

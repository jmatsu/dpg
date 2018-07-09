package api

import (
	"gopkg.in/guregu/null.v3"
	"testing"
)

func TestAppsEndpoint_ToURL(t *testing.T) {
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
		t.Logf("TestAppsEndpoint_ToURL at %d", i)

		if c.in.ToURL() != c.expected {
			t.Errorf("%s was expected but %s was found.", c.expected, c.in.ToURL())
		}
	}
}

func TestAppMembersEndpoint_ToURL(t *testing.T) {
	cases := []struct {
		in       AppMembersEndpoint
		expected string
	}{
		{
			in: AppMembersEndpoint{
				BaseURL:      "x",
				AppOwnerName: "y",
				AppPlatform:  "android",
				AppId:        "package",
			},
			expected: "x/api/users/y/platforms/android/apps/package/members",
		},
	}

	for i, c := range cases {
		t.Logf("TestAppMembersEndpoint_ToURL at %d", i)

		if c.in.ToURL() != c.expected {
			t.Errorf("%s was expected but %s was found.", c.expected, c.in.ToURL())
		}
	}
}

func TestOrganizationAppTeamsEndpoint_ToURL(t *testing.T) {
	cases := []struct {
		in       OrganizationAppTeamsEndpoint
		expected string
	}{
		{
			in: OrganizationAppTeamsEndpoint{
				BaseURL:          "x",
				OrganizationName: "y",
				AppPlatform:      "android",
				AppId:            "package",
			},
			expected: "x/api/organizations/y/platforms/android/apps/package/teams",
		},
		{
			in: OrganizationAppTeamsEndpoint{
				BaseURL:          "x",
				OrganizationName: "y",
				AppPlatform:      "android",
				AppId:            "package",
				TeamName:         "z",
			},
			expected: "x/api/organizations/y/platforms/android/apps/package/teams/z",
		},
	}

	for i, c := range cases {
		t.Logf("TestOrganizationAppTeamsEndpoint_ToURL at %d", i)

		if c.in.ToURL() != c.expected {
			t.Errorf("%s was expected but %s was found.", c.expected, c.in.ToURL())
		}
	}
}

func TestOrganizationAppSharedTeamsEndpoint_ToURL(t *testing.T) {
	cases := []struct {
		in       EnterpriseOrganizationAppSharedTeamsEndpoint
		expected string
	}{
		{
			in: EnterpriseOrganizationAppSharedTeamsEndpoint{
				BaseURL:          "x",
				OrganizationName: "y",
				AppPlatform:      "android",
				AppId:            "package",
			},
			expected: "x/api/organizations/y/platforms/android/apps/package/shared_teams",
		},
		{
			in: EnterpriseOrganizationAppSharedTeamsEndpoint{
				BaseURL:          "x",
				OrganizationName: "y",
				AppPlatform:      "android",
				AppId:            "package",
				SharedTeamName:   "z",
			},
			expected: "x/api/organizations/y/platforms/android/apps/package/shared_teams/z",
		},
	}

	for i, c := range cases {
		t.Logf("TestOrganizationAppSharedTeamsEndpoint_ToURL at %d", i)

		if c.in.ToURL() != c.expected {
			t.Errorf("%s was expected but %s was found.", c.expected, c.in.ToURL())
		}
	}
}

func TestDistributionsEndpoint_ToURL(t *testing.T) {
	cases := []struct {
		in       DistributionsEndpoint
		expected string
	}{
		{
			in: DistributionsEndpoint{
				BaseURL: "x",
			},
			expected: "x/api/distributions",
		},
		{
			in: DistributionsEndpoint{
				BaseURL:         "x",
				DistributionKey: "y",
			},
			expected: "x/api/distributions/y",
		},
	}

	for i, c := range cases {
		t.Logf("TestDistributionsEndpoint_ToURL at %d", i)

		if c.in.ToURL() != c.expected {
			t.Errorf("%s was expected but %s was found.", c.expected, c.in.ToURL())
		}
	}
}

func TestOrganizationsEndpoint_ToURL(t *testing.T) {
	cases := []struct {
		in       OrganizationsEndpoint
		expected string
	}{
		{
			in: OrganizationsEndpoint{
				BaseURL: "x",
			},
			expected: "x/api/organizations",
		},
		{
			in: OrganizationsEndpoint{
				BaseURL:          "x",
				OrganizationName: "y",
			},
			expected: "x/api/organizations/y",
		},
	}

	for i, c := range cases {
		t.Logf("TestOrganizationsEndpoint_ToURL at %d", i)

		if c.in.ToURL() != c.expected {
			t.Errorf("%s was expected but %s was found.", c.expected, c.in.ToURL())
		}
	}
}

func TestOrganizationMembersEndpoint_ToURL(t *testing.T) {
	cases := []struct {
		in       OrganizationMembersEndpoint
		expected string
	}{
		{
			in: OrganizationMembersEndpoint{
				BaseURL:          "x",
				OrganizationName: "y",
			},
			expected: "x/api/organizations/y/members",
		},
		{
			in: OrganizationMembersEndpoint{
				BaseURL:          "x",
				OrganizationName: "y",
				UserName:         null.StringFrom("z"),
				UserEmail:        null.StringFromPtr(nil),
			},
			expected: "x/api/organizations/y/members/z",
		},
		{
			in: OrganizationMembersEndpoint{
				BaseURL:          "x",
				OrganizationName: "y",
				UserName:         null.StringFromPtr(nil),
				UserEmail:        null.StringFrom("z@why"),
			},
			expected: "x/api/organizations/y/members/z%40why",
		},
	}

	for i, c := range cases {
		t.Logf("TestOrganizationMembersEndpoint_ToURL at %d", i)

		if c.in.ToURL() != c.expected {
			t.Errorf("%s was expected but %s was found.", c.expected, c.in.ToURL())
		}
	}
}

func TestEnterpriseMembersEndpoint_ToURL(t *testing.T) {
	cases := []struct {
		in       EnterpriseMembersEndpoint
		expected string
	}{
		{
			in: EnterpriseMembersEndpoint{
				BaseURL:        "x",
				EnterpriseName: "y",
			},
			expected: "x/api/enterprises/y/users",
		},
		{
			in: EnterpriseMembersEndpoint{
				BaseURL:        "x",
				EnterpriseName: "y",
				UserName:       "z",
			},
			expected: "x/api/enterprises/y/users/z",
		},
	}

	for i, c := range cases {
		t.Logf("TestEnterpriseMembersEndpoint_ToURL at %d", i)

		if c.in.ToURL() != c.expected {
			t.Errorf("%s was expected but %s was found.", c.expected, c.in.ToURL())
		}
	}
}

func TestEnterpriseOrganizationsMembersEndpoint_ToURL(t *testing.T) {
	cases := []struct {
		in       EnterpriseOrganizationsMembersEndpoint
		expected string
	}{
		{
			in: EnterpriseOrganizationsMembersEndpoint{
				BaseURL:          "x",
				EnterpriseName:   "y",
				OrganizationName: "z",
			},
			expected: "x/api/enterprises/y/organizations/z/users",
		},
		{
			in: EnterpriseOrganizationsMembersEndpoint{
				BaseURL:          "x",
				EnterpriseName:   "y",
				OrganizationName: "z",
				UserName:         "A",
			},
			expected: "x/api/enterprises/y/organizations/z/users/A",
		},
	}

	for i, c := range cases {
		t.Logf("TestEnterpriseOrganizationsMembersEndpoint_ToURL at %d", i)

		if c.in.ToURL() != c.expected {
			t.Errorf("%s was expected but %s was found.", c.expected, c.in.ToURL())
		}
	}
}

func TestEnterpriseSharedTeamsEndpoint_ToURL(t *testing.T) {
	cases := []struct {
		in       EnterpriseSharedTeamsEndpoint
		expected string
	}{
		{
			in: EnterpriseSharedTeamsEndpoint{
				BaseURL:        "x",
				EnterpriseName: "y",
			},
			expected: "x/api/enterprises/y/shared_teams",
		},
		{
			in: EnterpriseSharedTeamsEndpoint{
				BaseURL:        "x",
				EnterpriseName: "y",
				SharedTeamName: "z",
			},
			expected: "x/api/enterprises/y/shared_teams/z",
		},
	}

	for i, c := range cases {
		t.Logf("TestEnterpriseSharedTeamsEndpoint_ToURL at %d", i)

		if c.in.ToURL() != c.expected {
			t.Errorf("%s was expected but %s was found.", c.expected, c.in.ToURL())
		}
	}
}

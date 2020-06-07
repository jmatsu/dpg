package api

import (
	"github.com/jmatsu/dpg/request/apps/members"
)

func (c Client) AddMember(app UserApp, request members.AddRequest) (string, error) {
	if err := app.verify(); err != nil {
		return "", err
	}

	if err := request.Verify(); err != nil {
		return "", err
	}

	endpoint := AppMembersEndpoint{
		BaseURL:      c.baseURL,
		AppOwnerName: app.OwnerName,
		AppId:        app.Id,
		AppPlatform:  app.Platform,
	}

	if bytes, err := endpoint.MultiPartFormRequest(c.authorization, request); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}

func (c Client) ListMembers(app UserApp, request members.ListRequest) (string, error) {
	if err := app.verify(); err != nil {
		return "", err
	}

	if err := request.Verify(); err != nil {
		return "", err
	}

	endpoint := AppMembersEndpoint{
		BaseURL:      c.baseURL,
		AppOwnerName: app.OwnerName,
		AppId:        app.Id,
		AppPlatform:  app.Platform,
	}

	if bytes, err := endpoint.GetListRequest(c.authorization, request); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}

func (c Client) RemoveMember(app UserApp, request members.RemoveRequest) (string, error) {
	if err := app.verify(); err != nil {
		return "", err
	}

	if err := request.Verify(); err != nil {
		return "", err
	}

	endpoint := AppMembersEndpoint{
		BaseURL:      c.baseURL,
		AppOwnerName: app.OwnerName,
		AppId:        app.Id,
		AppPlatform:  app.Platform,
	}

	if bytes, err := endpoint.DeleteRequest(c.authorization, request); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}

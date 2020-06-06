package api

import (
	"github.com/jmatsu/dpg/request/apps/members/add"
	"github.com/jmatsu/dpg/request/apps/members/list"
	"github.com/jmatsu/dpg/request/apps/members/remove"
)

func (c Client) AddNewMember(app UserApp, request add.Request) (string, error) {
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

func (c Client) ListMembers(app UserApp, request list.Request) (string, error) {
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

func (c Client) RemoveMember(app UserApp, request remove.Request) (string, error) {
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

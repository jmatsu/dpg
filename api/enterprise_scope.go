package api

import "github.com/jmatsu/dpg/request/enterprises/members"

func (c Client) AddEnterpriseMember(enterprise Enterprise, request members.AddRequest) (string, error) {
	if err := enterprise.verify(); err != nil {
		return "", err
	}

	if err := request.Verify(); err != nil {
		return "", err
	}

	endpoint := EnterpriseMembersEndpoint{
		BaseURL:        c.baseURL,
		EnterpriseName: enterprise.Name,
	}

	if bytes, err := endpoint.MultiPartFormRequest(c.authorization, request); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}

func (c Client) ListEnterpriseMembers(enterprise Enterprise, request members.ListRequest) (string, error) {
	if err := enterprise.verify(); err != nil {
		return "", err
	}

	if err := request.Verify(); err != nil {
		return "", err
	}

	endpoint := EnterpriseMembersEndpoint{
		BaseURL:        c.baseURL,
		EnterpriseName: enterprise.Name,
	}

	if bytes, err := endpoint.GetListRequest(c.authorization, request); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}

func (c Client) RemoveEnterpriseMember(enterprise Enterprise, userName string, request members.RemoveRequest) (string, error) {
	if err := enterprise.verify(); err != nil {
		return "", err
	}

	if err := request.Verify(); err != nil {
		return "", err
	}

	endpoint := EnterpriseMembersEndpoint{
		BaseURL:        c.baseURL,
		EnterpriseName: enterprise.Name,
		UserName:       userName,
	}

	if bytes, err := endpoint.DeleteRequest(c.authorization, request); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}

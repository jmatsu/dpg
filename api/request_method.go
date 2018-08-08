package api

import (
	"errors"
	"fmt"
	"github.com/jmatsu/dpg/api/request"
	"github.com/jmatsu/dpg/api/response"
	"github.com/jmatsu/dpg/version"
	"github.com/sirupsen/logrus"
	"net/http"
)

func getRequest(e Endpoint, authorization *Authorization, requestParams request.Params) ([]byte, error) {
	query, err := request.ToQuery(requestParams)

	if err != nil {
		return nil, err
	}

	logrus.Debugf("query = %s\n", query)

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s?%s", e.ToURL(), query), nil)
	authorization.doAuthorize(req)

	if err != nil {
		return nil, err
	}

	resp, err := new(http.Client).Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bytes, errResp, err := response.FilterErrorResponse(*resp)

	if err != nil {
		return nil, err
	}

	if errResp != nil {
		return nil, errors.New(errResp.Message)
	}

	return bytes, nil
}

func deleteRequest(e Endpoint, authorization *Authorization, requestBody request.Body) ([]byte, error) {
	data, contentType, err := request.ToMultiFormPart(requestBody)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodDelete, e.ToURL(), &data)
	req.Header.Set("User-Agent", version.UserAgent())
	authorization.doAuthorize(req)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", contentType)

	resp, err := new(http.Client).Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bytes, errResp, err := response.FilterErrorResponse(*resp)

	if err != nil {
		return nil, err
	}

	if errResp != nil {
		return nil, errors.New(fmt.Sprintf("api returned an error response : %s", errResp.Message))
	}

	return bytes, nil
}

func multiPartFormRequest(e Endpoint, authorization *Authorization, requestBody request.Body) ([]byte, error) {
	data, contentType, err := request.ToMultiFormPart(requestBody)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, e.ToURL(), &data)
	authorization.doAuthorize(req)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", contentType)

	resp, err := new(http.Client).Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bytes, errResp, err := response.FilterErrorResponse(*resp)

	if err != nil {
		return nil, err
	}

	if errResp != nil {
		return nil, errors.New(fmt.Sprintf("api returned an error response : %s", errResp.Message))
	}

	return bytes, nil
}

func patchRequest(e Endpoint, authorization *Authorization, requestBody request.Body) ([]byte, error) {
	data, contentType, err := request.ToMultiFormPart(requestBody)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPatch, e.ToURL(), &data)
	authorization.doAuthorize(req)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", contentType)

	resp, err := new(http.Client).Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bytes, errResp, err := response.FilterErrorResponse(*resp)

	if err != nil {
		return nil, err
	}

	if errResp != nil {
		return nil, errors.New(fmt.Sprintf("api returned an error response : %s", errResp.Message))
	}

	return bytes, nil
}

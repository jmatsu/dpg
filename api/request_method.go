package api

import (
	"github.com/jmatsu/dpg/api/request"
	"net/http"
	"github.com/jmatsu/dpg/api/response"
	"errors"
	"fmt"
)

func getRequest(e Endpoint, authority Authority, requestParams request.Params, verbose bool) ([]byte, error) {
	stringMap, err := requestParams.StringMap()

	if err != nil {
		return nil, err
	}

	query, err := authority.GetParams(stringMap)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s?%s", e.ToURL(), query), nil)

	if err != nil {
		return nil, err
	}

	resp, err := new(http.Client).Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bytes, errResp, err := response.FilterErrorResponse(*resp, verbose)

	if err != nil {
		return nil, err
	}

	if errResp != nil {
		return nil, errors.New(fmt.Sprintf("api returned an error response : %s", errResp.Message))
	}

	return bytes, nil
}

func multiPartFormRequest(e Endpoint, authority Authority, requestBody request.Body, verbose bool) ([]byte, error) {
	ioMap, err := requestBody.IoReaderMap()

	if err != nil {
		return nil, err
	}

	data, contentType, err := authority.MultiPartForm(ioMap)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, e.ToURL(), &data)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", contentType)

	resp, err := new(http.Client).Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bytes, errResp, err := response.FilterErrorResponse(*resp, verbose)

	if err != nil {
		return nil, err
	}

	if errResp != nil {
		return nil, errors.New(fmt.Sprintf("api returned an error response : %s", errResp.Message))
	}

	return bytes, nil
}

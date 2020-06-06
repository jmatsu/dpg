package response

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/guregu/null.v3"
	"io/ioutil"
	"net/http"
	"os"
)

type ErrorResponse struct {
	IsError bool   `json:"error"`
	Message string `json:"message"`
}

type errorResponseMock struct {
	IsError bool        `json:"error"`
	Message null.String `json:"message"`
	Because null.String `json:"because"`
}

func (mock errorResponseMock) ensure() *ErrorResponse {
	return &ErrorResponse{
		IsError: mock.IsError,
		Message: mock.Message.String,
	}
}

func FilterErrorResponse(resp http.Response) (bytes []byte, errorResponse *ErrorResponse, failure error) {
	bytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		failure = err
		return
	}

	errResp := errorResponseMock{}

	if err := json.Unmarshal(bytes, &errResp); err != nil {
		if _, ok := err.(*json.SyntaxError); ok {
			response := string(bytes)

			if response == "" {
				logrus.Errorln("The response is no content. This seems not to be expected.")
				logrus.Errorln("Please create an issue to https://github.com/jmatsu/dpg/issues if possible. Thanks.")
			} else {
				logrus.Errorln(response)
			}
		}

		failure = err
		return
	}

	if code := resp.StatusCode; errResp.IsError {
		defer func() {
			logrus.Warnf("an error was returned with code %d\n", code)
		}()

		errorResponse = errResp.ensure()
	}

	defer func() {
		fmt.Fprintln(os.Stdout, string(bytes))
	}()

	return
}

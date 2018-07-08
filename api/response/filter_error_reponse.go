package response

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func FilterErrorResponse(resp http.Response) (bytes []byte, errorResponse *ErrorResponse, failure error) {
	bytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		failure = err
		return
	}

	errResp := ErrorResponseMock{}

	if err := json.Unmarshal(bytes, &errResp); err != nil {
		failure = err
		return
	}

	if code := resp.StatusCode; errResp.IsError {
		defer func() {
			fmt.Fprintf(os.Stderr, "an error was returned with code %d\n", code)
		}()

		errorResponse = errResp.Ensure()
	}

	defer func() {
		fmt.Fprintf(os.Stdout, "%s\n", string(bytes))
	}()

	return
}

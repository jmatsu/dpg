package api

import (
	"encoding/json"
	"fmt"
	"github.com/jmatsu/dpg/api/response"
	"github.com/jmatsu/dpg/util"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Authority struct {
	Token string
}

type Key string

const (
	Token Key = "token"
)

func (authority Authority) ioReaderMap() (*map[string]io.Reader, error) {
	m := map[Key]io.Reader{
		Token: strings.NewReader(authority.Token),
	}

	out, err := util.StringifyKeys(m)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (authority Authority) Merge(m map[string]io.Reader) (out map[string]io.Reader, err error) {
	m2, err := authority.ioReaderMap()

	if err != nil {
		return
	}

	out = util.Merge(m, *m2)

	return
}

func Response(resp http.Response, verbose bool) (bytes []byte, errorResponse *response.ErrorResponse, failure error) {
	bytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		failure = err
		return
	}

	if verbose {
		errResp := response.ErrorResponseMock{}

		if err := json.Unmarshal(bytes, &errResp); err != nil {
			failure = err
			return
		}

		if code := resp.StatusCode; errResp.IsError {
			defer func() {
				fmt.Fprintf(os.Stderr, "An error was returned with code %d\n", code)
			}()

			errorResponse = errResp.Ensure()
		}

		defer func() {
			fmt.Fprintf(os.Stdout, string(bytes))
		}()
	}

	return
}

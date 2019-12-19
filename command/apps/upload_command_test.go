package apps

import (
	"fmt"
	"github.com/jmatsu/dpg/api"
	requestAppUpload "github.com/jmatsu/dpg/api/request/apps/upload"
	"gopkg.in/guregu/null.v3"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

type RequestCase struct {
	command uploadCommand

	expectedResponse bool
	expectError      bool
}

func existingFilePath(fileNameExtension string) string {
	return fmt.Sprintf("%s/src/github.com/jmatsu/dpg/fixture/test.%s", os.Getenv("GOPATH"), fileNameExtension)
}

func createRequestCases(fileNameExtension string) []RequestCase {
    existingFilePath := existingFilePath(fileNameExtension)
    return []RequestCase {
	{
		command: uploadCommand{
			endpoint: &api.AppsEndpoint{
				BaseURL: "xxxx",
			},
			requestBody: &requestAppUpload.Request{
				AppFilePath:        existingFilePath,
				AppVisible:         false,
				EnableNotification: true,
				ShortMessage:       null.StringFrom("xxxxxx"),
				DistributionKey:    null.StringFrom("xxxxxx"),
				DistributionName:   null.StringFrom("xxxxxx"),
				ReleaseNote:        null.StringFrom("xxxxxx"),
			},
		},
		expectError: true,
	},
	{
		command: uploadCommand{
			endpoint: &api.AppsEndpoint{
				BaseURL:      "xxxx",
				AppOwnerName: "xx",
			},
			requestBody: &requestAppUpload.Request{
				AppVisible:         false,
				EnableNotification: true,
				ShortMessage:       null.StringFrom("xxxxxx"),
				DistributionKey:    null.StringFrom("xxxxxx"),
				DistributionName:   null.StringFrom("xxxxxx"),
				ReleaseNote:        null.StringFrom("xxxxxx"),
			},
		},
		expectError: true,
	},
	{
		command: uploadCommand{
			endpoint: &api.AppsEndpoint{
				BaseURL:      "xxxx",
				AppOwnerName: "xx",
			},
			requestBody: &requestAppUpload.Request{
				AppFilePath:        "not exists",
				AppVisible:         false,
				EnableNotification: true,
				ShortMessage:       null.StringFrom("xxxxxx"),
				DistributionKey:    null.StringFrom("xxxxxx"),
				DistributionName:   null.StringFrom("xxxxxx"),
				ReleaseNote:        null.StringFrom("xxxxxx"),
			},
		},
		expectError: true,
	},
	{
		command: uploadCommand{
			endpoint: &api.AppsEndpoint{
				BaseURL:      "xxxx",
				AppOwnerName: "xx",
			},
			requestBody: &requestAppUpload.Request{
				AppFilePath:        existingFilePath,
				AppVisible:         false,
				EnableNotification: true,
				ShortMessage:       null.StringFrom("xxxxxx"),
				DistributionKey:    null.StringFrom(""),
				DistributionName:   null.StringFrom("xxxxxx"),
				ReleaseNote:        null.StringFrom("xxxxxx"),
			},
		},
		expectError: true,
	},
	{
		command: uploadCommand{
			endpoint: &api.AppsEndpoint{
				BaseURL:      "xxxx",
				AppOwnerName: "xx",
			},
			requestBody: &requestAppUpload.Request{
				AppFilePath:        existingFilePath,
				EnableNotification: true,
				ShortMessage:       null.StringFrom("xxxxxx"),
				DistributionKey:    null.StringFrom("xxxxxx"),
				DistributionName:   null.StringFrom("xxxxxx"),
				ReleaseNote:        null.StringFrom("xxxxxx"),
			},
		},
		expectedResponse: true,
	},
	{
		command: uploadCommand{
			endpoint: &api.AppsEndpoint{
				BaseURL:      "xxxx",
				AppOwnerName: "xx",
			},
			requestBody: &requestAppUpload.Request{
				AppFilePath:        existingFilePath,
				AppVisible:         false,
				EnableNotification: true,
				ShortMessage:       null.StringFrom("xxxxxx"),
				DistributionKey:    null.StringFrom("xxxxxx"),
				DistributionName:   null.StringFrom(""),
				ReleaseNote:        null.StringFrom("xxxxxx"),
			},
		},
		expectedResponse: true,
	},
	{
		command: uploadCommand{
			endpoint: &api.AppsEndpoint{
				BaseURL:      "xxxx",
				AppOwnerName: "xx",
			},
			requestBody: &requestAppUpload.Request{
				AppFilePath:        existingFilePath,
				AppVisible:         true,
				EnableNotification: true,
				ShortMessage:       null.StringFrom("xxxxxx"),
				DistributionKey:    null.StringFrom("xxxxxx"),
				DistributionName:   null.StringFrom("xxxxxx"),
				ReleaseNote:        null.StringFrom("xxxxxx"),
			},
		},
		expectedResponse: true,
	},
	{
		command: uploadCommand{
			endpoint: &api.AppsEndpoint{
				BaseURL:      "xxxx",
				AppOwnerName: "xx",
			},
			requestBody: &requestAppUpload.Request{
				AppFilePath:        existingFilePath,
				AppVisible:         false,
				EnableNotification: true,
				ShortMessage:       null.StringFrom("xxxxxx"),
				DistributionKey:    null.StringFrom("xxxxxx"),
				DistributionName:   null.StringFrom("xxxxxx"),
				ReleaseNote:        null.StringFrom("xxxxxx"),
			},
		},
		expectedResponse: true,
	},
    }
}

func TestVerifyInput(t *testing.T) {
    requestTestCases := append(
        createRequestCases("apk"),
        createRequestCases("aab")...
    )
	for i, c := range requestTestCases {
		t.Logf("TestVerifyInput at %d", i)

		if err := c.command.VerifyInput(); err == nil && c.expectError {
			t.Error("an error was expected.")
		}
	}
}

func TestRun(t *testing.T) {
    requestTestCases := append(
        createRequestCases("apk"),
        createRequestCases("aab")...
    )
	for i, c := range requestTestCases {
		t.Logf("TestRun at %d", i)

		if c.expectError {
			continue
		}

		func() {
			apiStub := mockServer(t, c.expectedResponse, c.expectError)

			defer apiStub.Close()

			c.command.endpoint.BaseURL = apiStub.URL
			c.command.Run(&api.Authorization{
				Token: "token",
			})
		}()
	}
}

func mockServer(t *testing.T, expectedResponse bool, expectError bool) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if expectedResponse {
			if r.Header.Get("Authorization") != "Token token" {
				t.Fatal("authorized header is wrong")
			}

			if !expectError {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("{\"error\": false}"))
			} else {
				t.Fatal("an unexpected request happened")
			}
		} else {
			t.Error("no response was expected but a server got a response\n")
		}
	}))
}

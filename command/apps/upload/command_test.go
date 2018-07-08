package apps_upload

import (
	"fmt"
	"github.com/jmatsu/dpg/api"
	requestAppUpload "github.com/jmatsu/dpg/api/request/apps/upload"
	"gopkg.in/guregu/null.v3"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func existingFilePath() string {
	return fmt.Sprintf("%s/src/github.com/jmatsu/dpg/fixture/test.%s", os.Getenv("GOPATH"), "apk")
}

var verifyCases = []struct {
	endpoint    api.AppUploadEndpoint
	authority   api.Authority
	requestBody requestAppUpload.Request

	expectError bool
}{
	{
		expectError: true,
	},
	{
		endpoint: api.AppUploadEndpoint{
			BaseURL: "xxxx",
		},
		requestBody: requestAppUpload.Request{
			AppFilePath:        existingFilePath(),
			AppVisible:         false,
			EnableNotification: true,
			ShortMessage:       null.StringFrom("xxxxxx"),
			DistributionKey:    null.StringFrom("xxxxxx"),
			DistributionName:   null.StringFrom("xxxxxx"),
			ReleaseNote:        null.StringFrom("xxxxxx"),
		},
		expectError: true,
	},
	{
		endpoint: api.AppUploadEndpoint{
			BaseURL:      "xxxx",
			AppOwnerName: "xx",
		},
		requestBody: requestAppUpload.Request{
			AppFilePath:        existingFilePath(),
			AppVisible:         false,
			EnableNotification: true,
			ShortMessage:       null.StringFrom("xxxxxx"),
			DistributionKey:    null.StringFrom("xxxxxx"),
			DistributionName:   null.StringFrom("xxxxxx"),
			ReleaseNote:        null.StringFrom("xxxxxx"),
		},
		expectError: true,
	},
	{
		endpoint: api.AppUploadEndpoint{
			BaseURL:      "xxxx",
			AppOwnerName: "xx",
		},
		authority: api.Authority{
			Token: "xxxx",
		},
		requestBody: requestAppUpload.Request{
			AppVisible:         false,
			EnableNotification: true,
			ShortMessage:       null.StringFrom("xxxxxx"),
			DistributionKey:    null.StringFrom("xxxxxx"),
			DistributionName:   null.StringFrom("xxxxxx"),
			ReleaseNote:        null.StringFrom("xxxxxx"),
		},
		expectError: true,
	},
	{
		endpoint: api.AppUploadEndpoint{
			BaseURL:      "xxxx",
			AppOwnerName: "xx",
		},
		authority: api.Authority{
			Token: "xxxx",
		},
		requestBody: requestAppUpload.Request{
			AppFilePath:        "not exists",
			AppVisible:         false,
			EnableNotification: true,
			ShortMessage:       null.StringFrom("xxxxxx"),
			DistributionKey:    null.StringFrom("xxxxxx"),
			DistributionName:   null.StringFrom("xxxxxx"),
			ReleaseNote:        null.StringFrom("xxxxxx"),
		},
		expectError: true,
	},
	{
		endpoint: api.AppUploadEndpoint{
			BaseURL:      "xxxx",
			AppOwnerName: "xx",
		},
		authority: api.Authority{
			Token: "xxxx",
		},
		requestBody: requestAppUpload.Request{
			AppFilePath:        existingFilePath(),
			AppVisible:         false,
			EnableNotification: true,
			ShortMessage:       null.StringFrom("xxxxxx"),
			DistributionKey:    null.StringFrom(""),
			DistributionName:   null.StringFrom("xxxxxx"),
			ReleaseNote:        null.StringFrom("xxxxxx"),
		},
		expectError: true,
	},
	{
		endpoint: api.AppUploadEndpoint{
			BaseURL:      "xxxx",
			AppOwnerName: "xx",
		},
		authority: api.Authority{
			Token: "xxxx",
		},
		requestBody: requestAppUpload.Request{
			AppFilePath:        existingFilePath(),
			AppVisible:         false,
			EnableNotification: true,
			ShortMessage:       null.StringFrom("xxxxxx"),
			DistributionKey:    null.StringFrom("xxxxxx"),
			DistributionName:   null.StringFrom(""),
			ReleaseNote:        null.StringFrom("xxxxxx"),
		},
		expectError: true,
	},
	{
		endpoint: api.AppUploadEndpoint{
			BaseURL:      "xxxx",
			AppOwnerName: "xx",
		},
		authority: api.Authority{
			Token: "xxxx",
		},
		requestBody: requestAppUpload.Request{
			AppFilePath:        existingFilePath(),
			EnableNotification: true,
			ShortMessage:       null.StringFrom("xxxxxx"),
			DistributionKey:    null.StringFrom("xxxxxx"),
			DistributionName:   null.StringFrom("xxxxxx"),
			ReleaseNote:        null.StringFrom("xxxxxx"),
		},
		expectError: false,
	},
	{
		endpoint: api.AppUploadEndpoint{
			BaseURL:      "xxxx",
			AppOwnerName: "xx",
		},
		authority: api.Authority{
			Token: "xxxx",
		},
		requestBody: requestAppUpload.Request{
			AppFilePath:        existingFilePath(),
			AppVisible:         true,
			EnableNotification: true,
			ShortMessage:       null.StringFrom("xxxxxx"),
			DistributionKey:    null.StringFrom("xxxxxx"),
			DistributionName:   null.StringFrom("xxxxxx"),
			ReleaseNote:        null.StringFrom("xxxxxx"),
		},
		expectError: false,
	},
}

func TestVerifyInput(t *testing.T) {
	for i, c := range verifyCases {
		t.Logf("testVerifyInput at %d", i)

		if err := verifyInput(c.endpoint, c.authority, c.requestBody); err == nil && c.expectError {
			t.Error("an error was expected.")
		}
	}
}

var appTestCases = []struct {
	token    string
	appOwner string
	request  requestAppUpload.Request

	expectedResponse bool
	expectError      bool
	expectedPanic    bool
}{
	{
		token:    "token",
		appOwner: "appOwner",
		request: requestAppUpload.Request{
			AppFilePath: existingFilePath(),
		},
		expectedResponse: true,
		expectError:      false,
	},
}

func TestAction(t *testing.T) {
	for i, c := range appTestCases {
		t.Logf("testApp at %d", i)

		testAction(t, c.token, c.appOwner, c.request, c.expectedResponse, c.expectError, c.expectedPanic)
	}
}

func testAction(t *testing.T, token string, appOwner string, request requestAppUpload.Request, expectedResponse bool, expectError bool, expectedPanic bool) {
	apiStub := mockServer(t, expectedResponse, expectError)

	defer apiStub.Close()

	endpoint := api.AppUploadEndpoint{
		BaseURL:      apiStub.URL,
		AppOwnerName: appOwner,
	}

	authority := api.Authority{
		Token: token,
	}

	defer func() {
		err := recover()

		if expectedPanic && err == nil {
			t.Error("panic was expected but no panic")
		}
	}()

	if _, err := uploadApp(
		endpoint,
		authority,
		request,
	); err != nil {
		if !expectError {
			t.Log("No error was expected")
			t.Error(err.Error())
		}
	}
}

func mockServer(t *testing.T, expectedResponse bool, expectError bool) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if expectedResponse {
			if !expectError {
				w.WriteHeader(http.StatusOK)

				bytes, err := ioutil.ReadFile(fmt.Sprintf("%s/src/github.com/jmatsu/dpg/fixture/response/apps/upload/ok.json", os.Getenv("GOPATH")))

				if err != nil {
					t.Error(err.Error())
				}

				w.Write(bytes)
			} else {
				// 200 would be returned when failed...
				w.WriteHeader(http.StatusOK)

				bytes, err := ioutil.ReadFile(fmt.Sprintf("%s/src/github.com/jmatsu/dpg/fixture/response/apps/upload/ng.json", os.Getenv("GOPATH")))

				if err != nil {
					t.Error(err.Error())
				}

				w.Write(bytes)
			}
		} else {
			t.Error("no response was expected but a server got a response\n")
		}
	}))
}

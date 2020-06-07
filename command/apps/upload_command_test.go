package apps

import (
	"fmt"
	"github.com/jmatsu/dpg/api"
	"gopkg.in/guregu/null.v3"
	"gopkg.in/urfave/cli.v2"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

type RequestCase struct {
	token              string
	appOwnerName       string
	filePath           string
	visible            bool
	enableNotification bool
	shortMessage       null.String
	distributionKey    null.String
	distributionName   null.String
	releaseNote        null.String
	isIOS              bool
	isAndroid          bool
}

func existingFilePath(fileNameExtension string) string {
	return fmt.Sprintf("%s/src/github.com/jmatsu/dpg/fixture/test.%s", os.Getenv("GOPATH"), fileNameExtension)
}

func createGoldCases() []RequestCase {
	return []RequestCase{
		{
			token:        "token",
			appOwnerName: "owner",
			filePath:     existingFilePath("apk"),
		},
		{
			token:        "token",
			appOwnerName: "owner",
			filePath:     existingFilePath("ipa"),
		},
		{
			token:        "token",
			appOwnerName: "owner",
			filePath:     existingFilePath("apk"),
			isAndroid:    true,
		},
		{
			token:        "token",
			appOwnerName: "owner",
			filePath:     existingFilePath("ipa"),
			isIOS:        true,
		},
		{
			token:        "token",
			appOwnerName: "owner",
			filePath:     existingFilePath("apk"),
			shortMessage: null.StringFrom("dummy"),
		},
		{
			token:            "token",
			appOwnerName:     "owner",
			filePath:         existingFilePath("apk"),
			distributionName: null.StringFrom("dummy"),
		},
		{
			token:           "token",
			appOwnerName:    "owner",
			filePath:        existingFilePath("apk"),
			distributionKey: null.StringFrom("dummy"),
		},
		{
			token:            "token",
			appOwnerName:     "owner",
			filePath:         existingFilePath("apk"),
			distributionName: null.StringFrom("dummy1"),
			distributionKey:  null.StringFrom("dummy2"),
		},
	}
}

func createBadCases() []RequestCase {
	return []RequestCase{
		{
			token: "",
		},
		{
			token: "token",
		},
		{
			token:        "token",
			appOwnerName: "owner",
			filePath:     "not found",
		},
		{
			token:            "token",
			appOwnerName:     "owner",
			filePath:         existingFilePath("apk"),
			distributionName: null.StringFrom(""),
		},
		{
			token:           "token",
			appOwnerName:    "owner",
			filePath:        existingFilePath("apk"),
			distributionKey: null.StringFrom(""),
		},
		{
			token:        "token",
			appOwnerName: "owner",
			filePath:     existingFilePath("apk"),
			releaseNote:  null.StringFrom(""),
		},
		{
			token:        "token",
			appOwnerName: "owner",
			filePath:     existingFilePath("apk"),
			shortMessage: null.StringFrom(""),
		},
		{
			token:        "token",
			appOwnerName: "owner",
			filePath:     existingFilePath("apk"),
			isAndroid:    true,
			isIOS:        true,
		},
	}
}

func testUploadCommand(t *testing.T, index int, c RequestCase) error {
	t.Logf("TestRun at %d", index)

	app := &cli.App{}
	app.Name = "test"
	app.Commands = []*cli.Command{
		UploadCommand(),
	}

	args := []string{"upload"}

	if c.token != "" {
		args = append(args, "--token", c.token)
	}

	if c.filePath != "" {
		args = append(args, "--app", c.filePath)
	}

	if c.shortMessage.Valid {
		args = append(args, "--message", c.shortMessage.String)
	}

	if c.distributionName.Valid {
		args = append(args, "--distribution-name", c.distributionName.String)
	}

	if c.distributionKey.Valid {
		args = append(args, "--distribution-key", c.distributionKey.String)
	}

	if c.releaseNote.Valid {
		args = append(args, "--release-note", c.releaseNote.String)
	}

	if c.enableNotification {
		args = append(args, "--enable-notification")
	}

	if c.visible {
		args = append(args, "--public")
	}

	if c.isIOS {
		args = append(args, "--ios")
	}

	if c.isAndroid {
		args = append(args, "--android")
	}

	return app.Run(args)
}

func TestUploadCommand(t *testing.T) {
	apiStub := mockServer()

	defer apiStub.Close()

	api.EndpointURL = apiStub.URL

	goldCases := createGoldCases()

	for i, c := range goldCases {
		if err := testUploadCommand(t, i, c); err != nil {
			t.Error(err)
		}
	}

	delta := len(goldCases)

	for i, c := range createBadCases() {
		if err := testUploadCommand(t, i+delta, c); err == nil {
			t.Error("an error was expected")
		}
	}
}

func mockServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		//noinspection GoUnhandledErrorResult
		w.Write([]byte("{\"error\": false}"))
	}))
}

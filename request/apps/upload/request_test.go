package upload

import (
	"fmt"
	"gopkg.in/guregu/null.v3"
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"testing"
)

func existingFilePath() string {
	return fmt.Sprintf("%s/src/github.com/jmatsu/dpg/fixture/test.%s", os.Getenv("GOPATH"), "apk")
}

func TestBoolToVisibility(t *testing.T) {
	cases := []struct {
		in       bool
		expected string
	}{
		{
			in:       true,
			expected: "public",
		},
		{
			in:       false,
			expected: "private",
		},
	}

	for i, c := range cases {
		t.Logf("Testing FormatVisibility at %d", i)

		if actual := boolToVisibility(c.in); actual != c.expected {
			t.Fatalf("Expected %s but %s", c.expected, actual)
		}
	}
}

func TestRequest_IOReaderMap(t *testing.T) {
	cases := []struct {
		request     Request
		expectError bool
	}{
		{
			request: Request{
				AppVisible:  true,
				AppFilePath: existingFilePath(),
			},
			expectError: false,
		},
		{
			request: Request{
				AppVisible:         false,
				AppFilePath:        existingFilePath(),
				EnableNotification: true,
				ShortMessage:       null.StringFrom("short message"),
				ReleaseNote:        null.StringFrom("release note"),
				DistributionName:   null.StringFrom("distribution name"),
				DistributionKey:    null.StringFrom("distribution key"),
			},
			expectError: false,
		},
		{
			request: Request{
				AppFilePath:        existingFilePath(),
				EnableNotification: false,
			},
			expectError: false,
		},
		{
			request: Request{
				AppVisible:         true,
				AppFilePath:        existingFilePath(),
				EnableNotification: false,
				ShortMessage:       null.StringFromPtr(nil),
				ReleaseNote:        null.StringFromPtr(nil),
				DistributionName:   null.StringFromPtr(nil),
				DistributionKey:    null.StringFromPtr(nil),
			},
			expectError: false,
		},
		{
			request: Request{
				AppVisible:         true,
				AppFilePath:        "not exists",
				EnableNotification: false,
				ShortMessage:       null.StringFromPtr(nil),
				ReleaseNote:        null.StringFromPtr(nil),
				DistributionName:   null.StringFromPtr(nil),
				DistributionKey:    null.StringFromPtr(nil),
			},
			expectError: true,
		},
	}

	for i, c := range cases {
		t.Logf("Testing MultiPartForm at %d", i)

		ioMap, err := c.request.IoReaderMap()

		if err != nil {
			if !c.expectError {
				t.Error(err)
			}

			continue
		}

		for key, value := range *ioMap {
			switch key {
			case "visibility":
				assertKeyValue(key, &value, boolToVisibility(c.request.AppVisible), t)
			case "file":
				if value == nil {
					t.Errorf("%s is nil", key)
				}
			case "message":
				assertKeyValue(key, &value, c.request.ShortMessage, t)
			case "distribution_key":
				assertKeyValue(key, &value, c.request.DistributionKey, t)
			case "distribution_name":
				assertKeyValue(key, &value, c.request.DistributionName, t)
			case "release_note":
				assertKeyValue(key, &value, c.request.ReleaseNote, t)
			case "disable_notify":
				assertKeyValue(key, &value, !c.request.EnableNotification, t)
			default:
				t.Errorf("%s is not found", key)
			}
		}
	}
}

func assertKeyValue(key string, reader *io.Reader, expected interface{}, t *testing.T) {
	if reader == nil && &expected != nil {
		t.Fatalf("Expected %s but nil was given for key : %s", expected.(string), key)
		return
	}

	if reader == nil {
		return
	}

	bytes, err := ioutil.ReadAll(*reader)

	if err != nil {
		t.Errorf("ReadAll error: %v", err)
		return
	}

	actual := string(bytes)

	var expectedStr string

	switch expected.(type) {
	case string:
		expectedStr = expected.(string)
	case bool:
		expectedStr = strconv.FormatBool(expected.(bool))
	case null.String:
		expectedStr = expected.(null.String).String
	case null.Bool:
		expectedStr = strconv.FormatBool(expected.(null.Bool).Bool)
	case null.Int:
		expectedStr = strconv.FormatInt(expected.(null.Int).Int64, 64)
	case null.Float:
		expectedStr = strconv.FormatFloat(expected.(null.Float).Float64, 'f', 6, 64)
	default:
		expectedStr = fmt.Sprintf("stringify failed for %s", reflect.TypeOf(expected).String())
	}

	if actual != expectedStr {
		t.Errorf("Expected %s but %s for key : %s", expectedStr, actual, key)
	}
}

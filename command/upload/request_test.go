package upload

import (
	"fmt"
	"gopkg.in/guregu/null.v3"
	"io"
	"io/ioutil"
	"reflect"
	"strconv"
	"testing"
)

func TestFormatVisibility(t *testing.T) {
	cases := []struct {
		in       string
		expected appVisibility
	}{
		{
			in:       "public",
			expected: publicApp,
		},
		{
			in:       "private",
			expected: privateApp,
		},
		{
			in:       "xxxxx",
			expected: privateApp,
		},
	}

	for i, c := range cases {
		t.Logf("Testing FormatVisibility at %d", i)

		if actual := FormatVisibility(c.in); actual != c.expected {
			t.Fatalf("Expected %s but %s", c.expected, actual)
		}
	}
}

func TestRequest_IOReaderMap(t *testing.T) {
	cases := []Request{
		{
			AppVisibility: FormatVisibility("public"),
			AppFilePath:   "/Users/jmatsu/Workspace/private-usage/cloud-functions-for-app-notification/android-sample/pushee/build/outputs/apk/debug/pushee-debug.apk",
		},
		{
			AppVisibility:        FormatVisibility("xxxxx"),
			AppFilePath:          "/Users/jmatsu/Workspace/private-usage/cloud-functions-for-app-notification/android-sample/pushee/build/outputs/apk/debug/pushee-debug.apk",
			SuppressNotification: true,
			ShortMessage:         null.StringFrom("short message"),
			ReleaseNote:          null.StringFrom("release note"),
			DistributionName:     null.StringFrom("distribution name"),
			DistributionKey:      null.StringFrom("distribution key"),
		},
		{
			AppVisibility:        publicApp,
			AppFilePath:          "/Users/jmatsu/Workspace/private-usage/cloud-functions-for-app-notification/android-sample/pushee/build/outputs/apk/debug/pushee-debug.apk",
			SuppressNotification: false,
			ShortMessage:         null.StringFromPtr(nil),
			ReleaseNote:          null.StringFromPtr(nil),
			DistributionName:     null.StringFromPtr(nil),
			DistributionKey:      null.StringFromPtr(nil),
		},
	}

	for i, request := range cases {
		t.Logf("Testing MultiPartForm at %d", i)

		ioMap, err := request.ioReaderMap()

		if err != nil {
			t.Fatal(err)
		}

		for key, value := range *ioMap {
			switch key {
			case "visibility":
				assertKeyValue(key, &value, request.AppVisibility, t)
			case "file":
				if value == nil {
					t.Fatalf("%s is nil", key)
				}
			case "message":
				assertKeyValue(key, &value, request.ShortMessage, t)
			case "distribution_key":
				assertKeyValue(key, &value, request.DistributionKey, t)
			case "distribution_name":
				assertKeyValue(key, &value, request.DistributionName, t)
			case "release_note":
				assertKeyValue(key, &value, request.ReleaseNote, t)
			case "disable_notify":
				assertKeyValue(key, &value, request.SuppressNotification, t)
			default:
				t.Fatalf("%s is not found", key)
			}
		}
	}
}

func assertKeyValue(key string, reader *io.Reader, expected interface{}, t *testing.T) {
	if reader == nil && &expected != nil {
		t.Fatalf("Expected %s but nil was given for key : %s", expected.(string), key)
	}

	if reader == nil {
		return
	}

	bytes, err := ioutil.ReadAll(*reader)

	if err != nil {
		t.Fatalf("ReadAll error: %v", err)
	}

	actual := string(bytes)

	var expectedStr string

	switch expected.(type) {
	case string:
		expectedStr = expected.(string)
	case appVisibility:
		expectedStr = string(expected.(appVisibility))
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
		t.Fatalf("Expected %s but %s for key : %s", expectedStr, actual, key)
	}
}

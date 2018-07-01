package upload

import (
	"github.com/jmatsu/dpg/api"
	"testing"
)

func TestApp(t *testing.T) {
	authority := api.Authority{
		Token: "xxx",
	}

	defer func() {
		err := recover()

		if err != nil {
			panic(err)
		}
	}()

	file := "xxx"

	if _, err := uploadApp(
		authority,
		"xxx",
		Request{
			AppFilePath:          file,
			SuppressNotification: true,
			AppVisibility:        privateApp,
		},
		true,
	); err != nil {
		t.Fatal(err.Error())
	}
}

package request

import (
	"reflect"
	"strings"
	"testing"
)

type TestParams struct{}

func (p TestParams) StringMap() (*map[string]string, error) {
	return &map[string]string{
		"foo1": "bar1",
		"foo2": "&bar2",
	}, nil
}

var queryCases = []struct {
	in       TestParams
	expected []string
}{
	{
		in: TestParams{},
		expected: []string{
			"foo1=bar1",
			"foo2=%26bar2",
		},
	},
}

func TestToQuery(t *testing.T) {
	for i, c := range queryCases {
		t.Logf("TestToQuery at %d", i)

		if actual, err := ToQuery(c.in); err != nil {
			t.Error(err.Error())
		} else if !reflect.DeepEqual(strings.Split(actual, "&"), c.expected) {
			t.Errorf("%s was expected but %s\n", c.expected, actual)
		}
	}
}

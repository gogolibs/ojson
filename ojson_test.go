package ojson_test

import (
	"github.com/gogolibs/ojson"
	"github.com/stretchr/testify/require"
	"testing"
)

var goodCases = []struct {
	name     string
	expected string
	actual   ojson.Anything
}{
	{
		name:     "string: simple",
		expected: `"hello"`,
		actual:   "hello",
	},
	{
		name:     "object: simple",
		expected: `{"key":"value"}`,
		actual:   ojson.Object{"key": "value"},
	},
}

var badCases = []struct {
	name     string
	expected string
	actual   ojson.Anything
}{
	{
		name:     "bad JSON",
		expected: `failed to marshal map[string]interface {}{"field":(*chan string)`,
		actual:   ojson.Object{"field": new(chan string)},
	},
}

func TestMustMarshal(t *testing.T) {
	for _, testCase := range goodCases {
		t.Run(testCase.name, func(t *testing.T) {
			actualData := ojson.MustMarshal(testCase.actual)
			require.Equal(t, testCase.expected, string(actualData))
		})
	}
	for _, testCase := range badCases {
		t.Run(testCase.name, func(t *testing.T) {
			defer func() {
				r := recover()
				require.NotNil(t, r)
				require.Contains(t, r, testCase.expected)
			}()
			ojson.MustMarshal(testCase.actual)
		})
	}
}

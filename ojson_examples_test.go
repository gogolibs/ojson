package ojson_test

import (
	"github.com/gogolibs/ojson"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestObjectExample(t *testing.T) {
	obj := ojson.Object{
		"key1": "value1",
		"key2": "value2",
	}
	expected := `{"key1":"value1","key2":"value2"}`
	actual := string(ojson.MustMarshal(obj))
	require.Equal(t, expected, actual)
}

func TestArrayExample(t *testing.T) {
	arr := ojson.Array{"one", "two", "three"}
	expected := `["one","two","three"]`
	actual := string(ojson.MustMarshal(arr))
	require.Equal(t, expected, actual)
}

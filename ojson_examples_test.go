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

func TestMergeObjectsExample(t *testing.T) {
	obj1 := ojson.Object{
		"key1": "value1",
	}
	obj2 := ojson.Object{
		"key2": "value2",
	}
	obj := ojson.Merge(obj1, obj2)
	expected := `{"key1":"value1","key2":"value2"}`
	actual := string(ojson.MustMarshal(obj))
	require.Equal(t, expected, actual)
}

func TestConcatArraysExample(t *testing.T) {
	arr1 := ojson.Array{"one", "two"}
	arr2 := ojson.Array{"three", "four"}
	arr := ojson.Concat(arr1, arr2)
	expected := `["one","two","three","four"]`
	actual := string(ojson.MustMarshal(arr))
	require.Equal(t, expected, actual)
}

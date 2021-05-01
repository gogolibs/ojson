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

func TestCommonObjectExample(t *testing.T) {
	commonObj := ojson.Object{
		"key1": "value1",
	}
	obj1 := ojson.Object{
		"key2": "value2",
	}.Update(commonObj)
	obj2 := ojson.Object{
		"key3": "value3",
	}.Update(commonObj)
	expected1 := `{"key1":"value1","key2":"value2"}`
	actual1 := string(ojson.MustMarshal(obj1))
	require.Equal(t, expected1, actual1)
	expected2 := `{"key1":"value1","key3":"value3"}`
	actual2 := string(ojson.MustMarshal(obj2))
	require.Equal(t, expected2, actual2)
}

func TestCommonArrayExample(t *testing.T) {
	commonArr := ojson.Array{"five", "six"}
	arr1 := ojson.Array{"one", "two"}.Extend(commonArr)
	arr2 := ojson.Array{"three", "four"}.Extend(commonArr)
	expected1 := `["one","two","five","six"]`
	actual1 := string(ojson.MustMarshal(arr1))
	require.Equal(t, expected1, actual1)
	expected2 := `["three","four","five","six"]`
	actual2 := string(ojson.MustMarshal(arr2))
	require.Equal(t, expected2, actual2)
}

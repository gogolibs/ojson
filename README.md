# ojson #

[![GoDoc](https://godoc.org/github.com/gogolibs/ojson?status.svg)](https://pkg.go.dev/github.com/gogolibs/ojson)
[![Go Report Card](https://goreportcard.com/badge/github.com/gogolibs/ojson)](https://goreportcard.com/report/github.com/gogolibs/ojson)
[![CI](https://github.com/gogolibs/ojson/actions/workflows/test-and-coverage.yml/badge.svg)](https://github.com/gogolibs/ojson/actions/workflows/test-and-coverage.yml)
[![codecov](https://codecov.io/gh/gogolibs/ojson/branch/main/graph/badge.svg?token=JXSDP6Ifxi)](https://codecov.io/gh/gogolibs/ojson)

**ojson** is a collection of convenience type aliases to represent JSON objects in plain Go.

## Objects ##

```go
package example

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
```

## Arrays ##

```go
package example

import (
	"github.com/gogolibs/ojson"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestArrayExample(t *testing.T) {
	arr := ojson.Array{"one", "two", "three"}
	expected := `["one","two","three"]`
	actual := string(ojson.MustMarshal(arr))
	require.Equal(t, expected, actual)
}
```

## Common Object ##

```go
package example

import (
	"github.com/gogolibs/ojson"
	"github.com/stretchr/testify/require"
	"testing"
)

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
```

## Common Array ##

```go
package example

import (
	"github.com/gogolibs/ojson"
	"github.com/stretchr/testify/require"
	"testing"
)

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
```
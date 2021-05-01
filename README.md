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

## Merge Objects ##

```go
package example

import (
	"github.com/gogolibs/ojson"
	"github.com/stretchr/testify/require"
	"testing"
)

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
```

## Concatenate Arrays ##

```go
package example

import (
	"github.com/gogolibs/ojson"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestConcatArraysExample(t *testing.T) {
	arr1 := ojson.Array{"one", "two"}
	arr2 := ojson.Array{"three", "four"}
	arr := ojson.Concat(arr1, arr2)
	expected := `["one","two","three","four"]`
	actual := string(ojson.MustMarshal(arr))
	require.Equal(t, expected, actual)
}
```
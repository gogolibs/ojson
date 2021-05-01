package ojson

import (
	"encoding/json"
	"fmt"
)

// Anything is a type alias for an interface{}: it represents an any type in JSON.
type Anything = interface{}

// Any is a convenience alias for Anything.
type Any = Anything

// Object is a type alias for map[string]interface{}: it represents an object in JSON.
type Object = map[string]interface{}

// Obj is a convenience alias for Object.
type Obj = Object

// Array is a type alias for []interface{}: it represents an array in JSON.
type Array = []interface{}

// Arr is a convenience alias for Array.
type Arr = Array

// MustMarshal is a wrapper around json.Marshal, that panics in case of an error.
func MustMarshal(a Anything) []byte {
	data, err := json.Marshal(a)
	if err != nil {
		panic(fmt.Sprintf(`failed to marshal %#v to JSON`, a))
	}
	return data
}

// Merge merges multiple objects into single object
func Merge(objects ...Object) Object {
	result := Object{}
	for _, object := range objects {
		for key, value := range object {
			result[key] = value
		}
	}
	return result
}

// Concat concatenates multiple arrays into single array
func Concat(arrays ...Array) Array {
	n := 0
	for _, a := range arrays {
		n += len(a)
	}
	result := make(Array, 0, n)
	for _, a := range arrays {
		result = append(result, a...)
	}
	return result
}

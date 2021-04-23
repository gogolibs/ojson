package ojson

import (
	"encoding/json"
	"fmt"
)

// Anything is a type alias for an interface{}: it represents an any type in JSON
type Anything = interface{}
// Any is a convenience alias for Anything
type Any = Anything
// Object is a type alias for map[string]interface{}: it represents an object in JSON
type Object = map[string]interface{}
// Obj is a convenience alias for Object
type Obj = Object
// O is a convenience alias for Object
type O = Object
// Array is a type alias for []interface{}: it represents an array in JSON
type Array = []interface{}
// Arr is a convenience alias for Array
type Arr = Array
// A is a convenience alias for Array
type A = Arr

// MustMarshal is a wrapper around json.Marshal, that panics in case of an error
func MustMarshal(a Anything) []byte {
	data, err := json.Marshal(a)
	if err != nil {
		panic(fmt.Sprintf(`failed to marshal %#v to JSON`, a))
	}
	return data
}

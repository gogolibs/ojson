package ojson

import (
	"encoding/json"
	"fmt"
)

type Anything = interface{}
type Any = Anything
type Object = map[string]interface{}
type Obj = Object
type O = Object
type Array = []interface{}
type Arr = Array
type A = Arr

func MustMarshal(a Anything) []byte {
	data, err := json.Marshal(a)
	if err != nil {
		panic(fmt.Sprintf(`failed to marshal %#v to JSON`, a))
	}
	return data
}

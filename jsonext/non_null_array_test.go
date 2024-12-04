package jsonext

import (
	"encoding/json"
	"fmt"
)

func ExampleNonNullArray() {
	for _, arr := range []NonNullArray[string]{
		[]string{"a"},
		[]string{},
		nil,
	} {
		b, err := json.Marshal(arr)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(b))
	}

	// Output:
	// ["a"]
	// []
	// []
}

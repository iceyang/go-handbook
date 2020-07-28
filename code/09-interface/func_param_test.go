package main

import (
	"fmt"
	"testing"
)

func variableLengthParams(params ...interface{}) {
	fmt.Println(params...)

	// Should Check length first.
	// fmt.Println(params[0])
}

func TestVariableLengthParams(t *testing.T) {
	variableLengthParams()
	variableLengthParams(nil)
	variableLengthParams("hello")
}

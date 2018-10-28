package main

import (
	"fmt"
	"testing"
)

//bonus: not only fields of the receiver, but also the body
type zipCode string

func (r *zipCode) updateZip(replacement string) {
	*r = zipCode(replacement)
}

func TestZip(t *testing.T) {

	fmt.Println("TestReplacing:")

	var v1 zipCode = "fizz"

	v1.updateZip("buzz")
	fmt.Println(v1) //prints: {buzz}
}

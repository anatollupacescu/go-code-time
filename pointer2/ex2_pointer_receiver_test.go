package main

import (
	"fmt"
	"testing"
)

type container struct {
	val string
}

func (r *container) setVal(newVal string) {
	r.val = newVal
}

func TestPointerReceiver(t *testing.T) {
	v := container{val: "fizz"}

	v.setVal("buzz")
	fmt.Println(v) //prints: {buzz}
}

//bonus: not only fields of the receiver, but also the body
type name string

func (r *name) update(replacement string) {
	*r = name(replacement)
}

func TestReplacing(t *testing.T) {

	var v1 name = "fizz"
	fmt.Println(v1) //prints: {fizz}

	v1.update("buzz")
	fmt.Println(v1) //prints: {buzz}
}

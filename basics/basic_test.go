package main

import (
	"fmt"
	"testing"
)

type (
	s struct {
		val string
	}
	p struct {
		val string
	}
)

//gives you a copy of the original (t)
func (r s) change(s string) {
	r.val = s
}

func TestValueReceiver(t *testing.T) {

	v := s{val: "fizz"}
	v.change("buzz")

	fmt.Println(v) //prints: {fizz}
}

//pointer
func (r *p) change(s string) {
	r.val = s
}

func TestPointerReceiver(t *testing.T) {
	v := p{val: "fizz"}

	v.change("buzz")
	fmt.Println(v) //prints: {buzz}
}

//bonus: not only fields of the receiver, but also the body
type name string

func (r *name) update(replacement string) {
	*r = name(replacement)
}

func TestReplacing(t *testing.T) {

	fmt.Println("TestReplacing:")

	var v1 name = "fizz"

	v1.update("buzz")
	fmt.Println(v1) //prints: {buzz}
}

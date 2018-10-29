package main

import (
	"fmt"
	"testing"
)

type data struct {
	val string
}

//gives you a copy of the original (t)
func (d data) copyWith(newVal string) data {
	d.val = newVal
	return d
}

func TestValueReceiver(t *testing.T) {

	v := data{val: "fizz"}
	fmt.Printf("we got %v\n", v) //prints: {fizz}

	buzz := v.copyWith("buzz")
	fmt.Printf("we got %v\n", buzz)

	if v == buzz {
		t.Fail()
	}
}

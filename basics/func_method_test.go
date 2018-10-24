package main

import (
	"fmt"
	"testing"
	"unsafe"
)

/* function */

func thisIsAFunction(arg []string) int {
	return len(arg)
}

/* method */

type data struct{}

func (data) thisIsAMethod() {

}

/* what gets copied and what not */

type example struct {
	i int
	c *complex128
	s []string
}

func TestSizes(t *testing.T) {
	var c complex128 = 1
	p := &c
	fmt.Println(unsafe.Sizeof(c))		//16 bytes
	fmt.Println(unsafe.Sizeof(p))		//8	bytes
}
package main

import (
	"fmt"
	"math"
	"testing"
	"unsafe"
)

type bigType complex128

/* what gets copied and what not */
type example struct {
	i int
	c *bigType
	s []string
}

func TestSizes(t *testing.T) {
	var c bigType = 1
	var p *bigType = &c
	fmt.Printf("big type size: %d\n", unsafe.Sizeof(c)) //16 bytes
	fmt.Printf("pointer size: %d\n", unsafe.Sizeof(p))  //8	bytes

	b := bigType(complex(math.MaxFloat64, math.MaxFloat64))
	var e example = example{3, &b, []string{"seven"}}

	fmt.Printf("struct size: %d\n", unsafe.Sizeof(e))   //8	bytes
	fmt.Printf("pointer size: %d\n", unsafe.Sizeof(&e)) //8	bytes
}

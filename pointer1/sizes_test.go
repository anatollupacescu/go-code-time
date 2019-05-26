package pointer1

import (
	"fmt"
	"math"
	"testing"
	"unsafe"
)

type bigType complex128

func TestSizes1(t *testing.T) {

	var c bigType = 1
	var p *bigType = &c

	fmt.Printf("big type size: %d\n", unsafe.Sizeof(c)) //16 bytes
	fmt.Printf("pointer size: %d\n", unsafe.Sizeof(p))  //8	bytes
}

type example struct {
	i int
	c *bigType
	s []string
}

func TestSizes2(t *testing.T) {

	b := bigType(complex(math.MaxFloat64, math.MaxFloat64))
	e := example{3, &b, []string{"seven"}}

	fmt.Printf("struct size: %d\n", unsafe.Sizeof(e))   //40	bytes
	fmt.Printf("pointer size: %d\n", unsafe.Sizeof(&e)) //8	bytes
}

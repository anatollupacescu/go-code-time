package main

import (
	"fmt"
	"testing"
	"unsafe"
)

type fixedStringData [1024]byte

func TestSize(t *testing.T) {

	var b fixedStringData

	fmt.Printf("size of the buffer: %6v\n", unsafe.Sizeof(b))
	fmt.Printf("len of the buffer: %7v\n", len(b))

	var c []byte = b[:]

	fmt.Println()
	fmt.Printf("size of the slice: %6v\n", unsafe.Sizeof(c))
	fmt.Printf("len of the slice: %7v\n", len(c))

	s := struct {
		r string
		f fixedStringData
	}{"", b}

	fmt.Println()
	fmt.Printf("size of the struct (value): %v\n", unsafe.Sizeof(s))

	var p = &s
	fmt.Printf("size of the struct pointer: %v\n", unsafe.Sizeof(p))

	var i interface{}
	fmt.Println()
	fmt.Printf("size of the interface: %v\n", unsafe.Sizeof(i))
}
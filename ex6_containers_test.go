package main

import (
	"fmt"
	"testing"
)

func TestValueSlice(t *testing.T) {

	var a []int = []int{3}

	b := make([]int, 1)
	copy(b, a)

	fmt.Printf("initial b: %d\n", b[0])

	a[0] = 4
	fmt.Printf("updated a: %d\n", a[0])
	fmt.Printf("value of b: %d\n", b[0])
}

func TestPointerSlice(t *testing.T) {

	var i int = 3
	var a []*int = []*int{&i}

	b := make([]*int, 1)
	copy(b, a)

	fmt.Printf("initial b: %d\n", *b[0])

	i = i + 1
	a[0] = &i
	fmt.Printf("updated a: %d\n", *a[0])
	fmt.Printf("value of b: %d\n", *b[0])
}

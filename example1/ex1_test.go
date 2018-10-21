package main

import (
	"fmt"
	"testing"
	"unsafe"
)

type s struct {
	name string
	data []byte
}

func (s *s) setName(n string) s {
	s.name = n
	return *s
}

func Test(t *testing.T) {
	var a s
	size := unsafe.Sizeof(a)
	fmt.Println(size)

	var b *s = &a
	size = unsafe.Sizeof(b)
	fmt.Println(size)
}

func fest(t *testing.T) {

	s := new(s)
	k := s.setName("test")

	fmt.Printf("k has name: %s\n", k.name)

	s2 := *s
	s2 = s2.setName("s2")

	fmt.Printf("s has name: %s\n", s.name)
	fmt.Printf("s2 has name: %s\n", s2.name)
}

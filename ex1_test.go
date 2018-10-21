package main

import (
	"fmt"
	"testing"
)

type s struct {
	name string
}

func (s *s) setName(n string) s {
	s.name = n
	return *s
}

func Test1(t *testing.T) {
	s := new(s)
	k := s.setName("test")

	fmt.Printf("k has name: %s\n", k.name)

	s2 := *s
	s2 = s2.setName("s2")

	fmt.Printf("s has name: %s\n", s.name)
	fmt.Printf("s2 has name: %s\n", s2.name)
}

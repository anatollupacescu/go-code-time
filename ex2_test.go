package main

import (
	"fmt"
	"testing"
)

type m struct {
	name string
}

func (m m) setName(n string) m {
	m.name = n
	return m
}

func (*m) test() {
	fmt.Print("lol\n")
}

func Test2(t *testing.T) {

	m := new(m)
	m2 := m.setName("jora")

	fmt.Printf("m has name: %s\n", m.name)
	fmt.Printf("m2 has name: %s\n", m2.name)
	m.test()
}

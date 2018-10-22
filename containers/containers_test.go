package main

import (
	"fmt"
	"testing"
	"unsafe"
)

type (

	userName [128]byte

	user struct {
		name userName
	}

	valueMap map[int]user

	pointerMap map[int]*user
)

func processValueMap(data valueMap) {
	fmt.Printf("size of value map: %v bytes\n", unsafe.Sizeof(data))
	data[0] = newDummyUser("value")
}

func processPointerMap(data pointerMap) {
	fmt.Printf("size of pointer map: %v bytes\n", unsafe.Sizeof(data))
	s := newDummyUser("pointer")
	data[0] = &s
}

func Test(t *testing.T) {
	s := "initial"

	u := newDummyUser(s)

	fmt.Printf("size of the user: %v\n", unsafe.Sizeof(u))

	valueMap := map[int]user{0: u}

	pointerMap := map[int]*user{0: &u}

	processValueMap(valueMap)
	assertNotEqual(valueMap[0], u, t)

	processPointerMap(pointerMap)
	assertNotEqual(*pointerMap[0], u, t)
}

func assertNotEqual(u1 user, u2 user, t *testing.T) {
	if u1 == u2 {
		t.Fatal("Users are equal")
	}
}

func newDummyUser(s string) user {
	var name userName
	copy(name[:], s)
	return user{name: name}
}

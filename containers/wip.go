package main

import "testing"

type slice []int

func (slice) thisIsFine() {
	//slice is a reference type, using pointers under-the-hood
}

func Test3(t *testing.T) {

}

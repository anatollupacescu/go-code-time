package main

import (
	"testing"
)

type slice []int

func (slice) thisIsFine() {
	//slice is a reference type, using pointers under-the-hood
}

/* if it's not changing the state, should it be a method in the first place? */
func maybeThisIsBetter(_ slice) {
}

func TestThis(t *testing.T) {
	var s slice

	s.thisIsFine()

	/*package.*/
	maybeThisIsBetter(s)
}

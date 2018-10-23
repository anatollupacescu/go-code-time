package main

type slice []int

func (slice) thisIsFine() {
	//slice is a reference type, using pointers under-the-hood
}

func maybeThisIsBetter(_ slice) {
}

func Test() {
	var s slice

	s.thisIsFine()

	/*package.*/maybeThisIsBetter(s)
}
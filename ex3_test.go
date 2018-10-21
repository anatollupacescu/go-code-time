package main

import (
	"fmt"
	"testing"
	"unsafe"
)

type movie struct {
	contents []byte //big
}

func findBestMovie(m []movie) movie {
	return movie{}
}

func Test3(t *testing.T) {

	movies := []movie{ /*lots of them*/ }
	findBestMovie(movies)
	fmt.Printf("size of movies: %v\n", unsafe.Sizeof(movies))
	many := []*movie{}
	fmt.Printf("size of many: %v\n", unsafe.Sizeof(many))
}

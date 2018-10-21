package main

import (
	"fmt"
	"testing"
)

type movieClip struct {
	contents   []byte //big
	actorCount int
}

func updateActorCount(clips []*movieClip) {
	for _, movie := range clips {
		movie.actorCount = 1 //calculated value
	}
}

func Test4(t *testing.T) {

	movies := []*movieClip{&movieClip{[]byte{}, 0} /*lots of them*/}
	updateActorCount(movies)

	fmt.Printf("actor count: %v\n", movies[0].actorCount)
}

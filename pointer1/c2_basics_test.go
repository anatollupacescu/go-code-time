package pointer1

import (
	"fmt"
	"testing"
)

/**
using pointers we can mutate a certain memory location cheaply (memory wise)
*/
func Test101(t *testing.T) {
	var f float64

	fc := f

	fc++

	fmt.Printf("got fc=%v and original=%v\n", fc, f)

	fp := &f

	*fp++

	fmt.Printf("got fp=%v and original=%v\n", *fp, f)
}

/**
some language constructs will make a copy under-the-hood
*/
func TestForLoop(t *testing.T) {

	values := []int{1, 2, 3, 4}

	for val := range values {
		fmt.Println(val)
		go func() {
			fmt.Println("inside", val)
		}()
	}
}

/**
a function call will make a copy of the arguments
FYI methods are a particular case of functions and same rules apply
*/
func TestFunctionCall(t *testing.T) {

	tt := complex(1, 2)

	func(tf complex128) {
		fmt.Printf("original addr=%v\n in func addr=%v\n", &tt, &tf)
	}(tt)

	fmt.Println()

	func(tf *complex128) {
		fmt.Printf("original addr=%v\n in func addr=%v\n", &tt, tf)
	}(&tt)
}

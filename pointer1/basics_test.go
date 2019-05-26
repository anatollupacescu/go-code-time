package pointer1

import (
	"fmt"
	"testing"
)

func Test101(t *testing.T) {
	var f float64

	fc := f

	fc++

	fmt.Printf("got fc=%v and original=%v\n", fc, f)

	fp := &f

	*fp++

	fmt.Printf("got fp=%v and original=%v\n", *fp, f)
}

func TestForLoop(t *testing.T) {

	arr := []int{6, 7}

	for i, v := range arr {
		fmt.Printf("got arr[%d]=%v at addr=%v and v at addr=%v\n", i, arr[i], &arr[i], &v)
	}

	//the danger!
}

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

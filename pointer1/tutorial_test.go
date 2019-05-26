package pointer1

import (
	"fmt"
	"testing"
)

type walletAmmount int

func (wa *walletAmmount) add(i int) {
	*wa += walletAmmount(i)
}

func TestBasicsPointer(t *testing.T) {

	var wa walletAmmount

	wa.add(100)

	fmt.Println(wa)

	copy := wa

	copy.add(50)

	fmt.Printf("original=%v, copy=%v\n", wa, copy)

	ptr := &wa

	ptr.add(99)

	fmt.Printf("original=%v, ptr=%v\n", wa, *ptr)
}

func TestForLoop(t *testing.T) {

	arr := []int{6, 7}

	for i, v := range arr {
		fmt.Printf("got arr[%d]=%v at addr=%v and v at addr=%v\n", i, arr[i], &arr[i], &v)
	}
}

type t struct {
	i float32
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

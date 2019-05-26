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

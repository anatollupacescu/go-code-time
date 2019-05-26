package pointer1

import (
	"fmt"
	"testing"
)

type walletAmmount int

func (wa *walletAmmount) add(i int) {
	*wa += walletAmmount(i)
}

func TestBasicsPointerReceiverUpdatesOriginal(t *testing.T) {

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

func (wa walletAmmount) addUseless(i int) {
	wa += walletAmmount(i)
}

func TestValueReceiverMakesACopy(t *testing.T) {

	var wa walletAmmount

	wa.addUseless(100)

	if wa == 0 {
		t.Fatal()
	}
}

//Bonus

func (wa walletAmmount) computeInterest(i float32) float64 {
	return 0.0
}

func computeInterest(wa walletAmmount, i float32) float64 {
	return 0.0
}

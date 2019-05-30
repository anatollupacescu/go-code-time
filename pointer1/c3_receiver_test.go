package pointer1

import (
	"fmt"
	"testing"
)

type walletAmount int

func (wa *walletAmount) add(i int) {
	*wa += walletAmount(i)
}

func (wa walletAmount) addUseless(i int) {
	wa += walletAmount(i)
}

func TestBasicsPointerReceiverUpdatesOriginal(t *testing.T) {

	var wa walletAmount

	wa.add(100)

	fmt.Println(wa)

	copy := wa

	copy.add(50)

	fmt.Printf("original=%v, copy=%v\n", wa, copy)

	ptr := &wa

	ptr.add(99)

	fmt.Printf("original=%v, ptr=%v\n", wa, *ptr)
}

/**
TODO: work around the compiler
*/
func TestValueReceiverMakesACopy(t *testing.T) {

	var wa walletAmount

	ptr := &wa

	ptr.addUseless(100)

	if wa == 0 {
		t.Fatal()
	}
}

/**
if a method does not mutate receiver, should it be a method
*/
func (wa walletAmount) computeInterest(i float32) float64 {
	return 0.0
}

func computeInterest(wa walletAmount, i float32) float64 {
	return 0.0
}

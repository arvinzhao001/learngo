package basic

import (
	"fmt"
	"testing"
)

// the control comes out of the switch statement immediately after a case is executed
func TestSwitch1(t *testing.T) {
	finger := 3
	switch finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Pinky")
	case 6, 7, 8, 9, 10: //multiple value
		fmt.Println("Use 1-5 number")
	default: //default case
		fmt.Println("incorrect finger number")
	}
}

//a fallthrough statement is used to transfer control to the first statement of the case
//that is present immediately after the case which has been executed
func TestSwitch2(t *testing.T) {
	switch num := 10; { //num is not a constant
	case num < 50:
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num < 5: //even through the condition is false, case block is executed
		fmt.Printf("%d is lesser than 5\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("%d is lesser than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d is lesser than 200\n", num)
	}
}

package basic

import (
	"fmt"
	"runtime/debug"
	"testing"
)

//1. There are some situations where the program cannot continue execution
//after a abnomal condition. In this case we use panic to prematurely terminate
//the program

func fullName(firstName *string, lastName *string) {
	defer recoverEx()
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

//recover is useful only when called inside deferred functions
//execute a call to recover inside a deferred function stops
//the panicking sequence by restoring normal execution.

//In this example, TestPanicCall not be effected by recoverEx
func recoverEx() {
	fmt.Println("deferred call in fullName")
	//r is the error message passed to the panic function
	if r := recover(); r != nil {
		fmt.Println("I don't want this exception effect the caller, recovered from ", r)
		debug.PrintStack()
	}
}

//1. When the program panics, any deferred function is executed and
//then the control returns to the caller whose deffered calls are
//executed and so on until the top level caller is reached.
//2. hence the program print the panic messages followed by the stack
//trace and then terminates
func TestPanicCall(t *testing.T) {
	defer fmt.Println("deferred call in main")
	firstName := "Elon"
	fullName(&firstName, nil)
	fmt.Println("returned normally from main")
}

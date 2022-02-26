package basic

import (
	"fmt"
	"testing"
)

//first class functions allows functions to be assigned to variables
//passed as arguments to other functions and returned from other functions

func TestAnonymousFunc1(t *testing.T) {
	a := func() {
		fmt.Println("test anonymous function")
	}
	a()
	fmt.Printf("%T\n", a)
}

func TestAnonymousFunc2(t *testing.T) {
	func(name string) {
		fmt.Println("Hello ", name)
	}("arvinzhao")
}

type op func(a int, b int) int

func TestDefineFuncType(t *testing.T) {
	var a op = func(a int, b int) int {
		return a + b
	}
	s := a(5, 6)
	fmt.Println("Sum", s)
}

//Pass functions as arguments to other functions
func TestHighOrderFunc1(t *testing.T) {
	m := func(a int, b int) int {
		return a * b
	}
	fmt.Println("Multi:", simple(m))
}

func simple(a op) int {
	return a(1, 3)
}

//Return functions from other functions
func TestHighOrderFunc2(t *testing.T) {
	m := addTwoNum()
	fmt.Println("Sum", m(2, 3))
}
func addTwoNum() func(int, int) int {
	f := func(a int, b int) int {
		return a + b
	}
	return f
}

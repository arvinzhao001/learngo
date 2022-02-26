package basic

import (
	"fmt"
	"testing"
)

//All the three components namely initialisation, condition and post are optional in Go
func TestFor1(t *testing.T) {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		if i == 5 {
			break
		}
		fmt.Printf("%d ", i)
	}
}

/*
labels can be used to break the outer loop from inside the inner for loop.
*/
func TestFor2(t *testing.T) {
outer:
	for i := 0; i < 3; i++ {
		for j := 1; j < 4; j++ {
			fmt.Printf("i = %d , j = %d\n", i, j)
			if i == j {
				break outer
			}
		}

	}
}

func TestFor3(t *testing.T) {
	i := 0
	for i <= 10 { // initialisation and post are omitted
		fmt.Printf("%d ", i)
		i += 2
	}
}

func TestFor4(t *testing.T) {
	i := 0
	for i <= 10 { //semicolons are ommitted and only condition is present
		fmt.Printf("%d ", i)
		i += 2
	}
}

func TestFor5(t *testing.T) {
	i := 1
	for { // infinite loop
		if i > 10 {
			break
		}
		fmt.Println("Hello World")
		i = i + 1
	}
}

func TestFor6(t *testing.T) {
	for no, i := 10, 1; i <= 10 && no <= 19; i, no = i+1, no+1 { //multiple initialisation and increment
		fmt.Printf("%d * %d = %d\n", no, i, no*i)
	}
}

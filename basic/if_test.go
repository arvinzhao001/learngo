package basic

import (
	"fmt"
	"testing"
)

/*
1. num is available only for access from inside the if and else
2. the else statement should start in the same line after the closing curly brace } of the if statement
*/
func TestIfAndVarible(t *testing.T) {
	if num := 10; num%2 == 0 {
		fmt.Println(num, " is even")
	} else {
		fmt.Println(num, " is odd")
	}
}

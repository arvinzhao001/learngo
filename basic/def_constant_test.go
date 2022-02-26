package basic

import (
	"fmt"
	"testing"
)

func TestDefineConstant(t *testing.T) {
	//declare a constant
	const num = 50
	//declare a group of constant
	const (
		str  = "test str1"
		fnum = 10.5
	)

	//Constants can not be reassigned again to any other value
	//num = 100

	//The value of constant should be known at the compile time
	//const num2 = math.Sqrt(4)

	//if not specify the type, constants is untyped, they supply type if and
	//only if a line of code demands it
	fmt.Printf("Type of num1:%T, Type of str1:%T, Type of fnum:%T", num, str, fnum)
	var intVar int = num
	var floatVar float32 = num
	var complexVar complex64 = num
	fmt.Println("intVar:", intVar, " floatVar:", floatVar, " complexVar:", complexVar)

}

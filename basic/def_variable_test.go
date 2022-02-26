package basic

import (
	"fmt"
	"testing"
)

func TestDefineVariable(t *testing.T) {
	/*If a variable is not assigned any value, Go automatically initialized it with zero value
	of the variable's type
	*/
	var age int
	fmt.Println("default value of age:", age)
	age = 32
	fmt.Println("value of age:", age)

	/*
		declaring a variable with a initial value
	*/
	var age2 int = 32
	fmt.Println("value of age2:", age2)

	/*
		if a variable has a inital value, Go will automatically be able to infer the type of the variable
		using that initial value
	*/
	var age3 = 32
	fmt.Println("value of age3:", age3)

	//short hand declaration
	age4 := 32
	fmt.Println("value of age4:", age4)

	//multiple variable declaration,
	var width, height int = 100, 200
	//the type can be removed if the variables have an inital value
	var width2, height2 = 300, 400
	//short hand syntax can only be used when at least one of the variables on the left side of := is newly declared
	width3, height3 := 500, 600
	fmt.Println("width is ", width, " height is ", height)
	fmt.Println("width2 is ", width2, " height2 is ", height2)
	fmt.Println("width3 is ", width3, " height3 is ", height3)

	type myInt int
	var oInt int = 10
	var nInt myInt = 20
	//Go is a strongly typed language,mixing types is not allowed, even we know myInt is alias of int
	//nInt = oInt
	fmt.Printf("Type of oInt: %T, Type of nInt: %T", oInt, nInt)

}

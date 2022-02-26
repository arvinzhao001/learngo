package basic

import (
	"fmt"
	"testing"
)

// a pointer is a variable that stores the memory address of another varible
func TestPointerBase(t *testing.T) {
	//method 1: declare a point from a variable
	base := 255
	var p1 *int = &base
	p2 := &base
	//address and dereference a pointer
	fmt.Println("the address of b is", p1, " ", p2)
	fmt.Println("the value of b is ", *p1, " ", *p2)

	//the zero value of pointer is nil
	var p3 *int
	if p3 == nil {
		p3 = &base
		fmt.Println("p3 is not initailized")
	}

	//method 2: declare a point directly
	size := new(int)
	fmt.Printf("Size value is %d, type is %T, address is %v\n", *size, size, size)
	*size = 85
	fmt.Println("New size value is ", *size)
	*size++
	fmt.Println("New size value is ", *size)

	//Go does not support pointer arthmetic
	//size++
}

func TestPassPointerToFunction(t *testing.T) {
	//change value
	val := 10
	singleValue(&val)
	fmt.Println("the value of val is ", val)

	//return pointer. the compiler does a escape analysis and
	//allocate val on the heap as the address escapes the local scope
	var val2 *int = returnPointer()
	fmt.Println("the value of val2 is ", val2)

	//Do not recommand to pass a pointer to array as an argument to function for modifying element
	arr1 := [...]int{10, 12, 14, 16}
	passArrayPointer1(&arr1)
	fmt.Println("after passArrayPointer1: ", arr1)

	//recommand to use slice instead
	passArrayPointer2(arr1[:])
	fmt.Println("after passArrayPointer2: ", arr1)
}

func singleValue(val *int) {
	*val = 100
}

func returnPointer() *int {
	val := 90
	return &val
}

//the argment is a pointer to array, type is *[4]int
func passArrayPointer1(arr *[4]int) {
	(*arr)[0] = 20
	arr[1] = 22 //short hand for (*arr)[1] = 22
}

func passArrayPointer2(slc []int) {
	slc[0] = 40
	slc[1] = 42
}

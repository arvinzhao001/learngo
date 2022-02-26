package basic

import (
	"fmt"
	"testing"
)

// array size is fixed
func TestDefArray(t *testing.T) {
	//declare an integer array of length 3, all elememt in an array are
	//automatically assigned the zero value of the array type
	var array1 [3]int
	array1[0] = 1
	fmt.Println(array1)

	//short hand declaration, and not necessary that all elements in an array have to be
	//assigned a value
	array2 := [3]int{1, 2}
	fmt.Println(array2)

	//ignore the length of the array, replace it with ...
	//and let the compiler find the length for you
	array4 := [...]int{7, 8, 9}

	//The size of the array is a part of the type, array cannot be resized
	/*
		var array5 int[5]
		array5 = array4  //assign failed, int[5] and int[3] is different type
	*/

	// a copy of array4 is assigned to array5.
	// Arrays in Go are value types and not reference types. when they are assigned to the new
	// variable or pass to a function. a copy of the original array is assigned to the new variale
	array6 := array4
	array6[0] = 10
	fmt.Println(array4)
	fmt.Println(array6)
}

//Passed by value, original array is unchanged
func TestTransferValue(t *testing.T) {
	a := [5]int{10, 11, 12, 13, 14}
	fmt.Println("Original:", a)
	transferValue(a)
	fmt.Println("After modify:", a)

}

func transferValue(arr [5]int) {
	arr[1] = arr[1] + 1
	fmt.Println("After modify", arr)
}

/*
Two method iterating array
*/
func TestIteArray(t *testing.T) {
	a := [...]float64{67.7, 11.1, 24.1, 88.9}

	fmt.Println("Iterating Method 1")
	for i := 0; i < len(a); i++ {
		fmt.Println("Index is ", i, " Value is ", a[i])
	}

	fmt.Println("Iterating Method 2")
	for i, v := range a {
		fmt.Println("Index is ", i, " Value is ", v)
	}
}

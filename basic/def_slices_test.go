package basic

import (
	"fmt"
	"testing"
)

//A slice with elements of type T is represented by []T
//slices do not own any data on their own, they are just references to existing arrays
func TestSliceBase(t *testing.T) {
	//method 1, create a slice from a array
	a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1:3]
	fmt.Println(b)
	fmt.Printf("Length of slice: %d, Capibility of slice: %d\n", len(b), cap(b))

	//method 2, create a slice
	c := []int{6, 7, 8}
	fmt.Println(c)

	//method3, use make(type,len,cap)
	//creates an array and return a slice reference to it
	d := make([]int, 5, 10)
	d[0] = 10
	fmt.Println(d)

	var e []int //the zero value of a slice type is nil
	fmt.Println(e == nil)
	e = append(e, 1)
	fmt.Printf("Length of slice: %d, Capibility of slice: %d\n", len(e), cap(e))
}

func TestSliceDynamic(t *testing.T) {
	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Println("Length of slice, ", len(slice1), "Capibility of slice", cap(slice1))
	/*
		when new elements are appended to the slice, a new array is created.
		The element of existing array is copied to the new array and a new slice
		reference for this  new array is returned. The capacity of the new slice
		is now twice of the old slice
	*/
	slice1 = append(slice1, 6)
	fmt.Println("Length of slice, ", len(slice1), "Capibility of slice", cap(slice1))

	//append one slice to another
	slice2 := []int{7, 8, 9}
	slice := append(slice1, slice2...)
	fmt.Println(slice)
	fmt.Println("Length of slice, ", len(slice1), "Capibility of slice", cap(slice1))
}

//when a slice is passed to a function as parameter, changes made
//inside the function are visible outside the function too
func TestSliceFunction(t *testing.T) {
	nos := []int{8, 7, 6}
	fmt.Println("slice before function call", nos)
	subtactOne(nos)                               //function modifies the slice
	fmt.Println("slice after function call", nos) //modifications are visible outside
}

func subtactOne(numbers []int) {
	for i := range numbers {
		numbers[i] -= 2
	}
}

/*
Slice hold a reference to the underlying array, As long as the slice is in memory.
the array cannot be garbage collected.
Use the copy function to make a copy of the slice. This way we can use the new slice.
and the original array can be garbage collected
*/
func TestOptGarbageCollect(t *testing.T) {
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	neededCountries := countries[:len(countries)-2]
	countriesCpy := make([]string, len(neededCountries))
	//copies neededCountries to countriesCpy, the original array can be garbage collected
	copy(countriesCpy, neededCountries)
}

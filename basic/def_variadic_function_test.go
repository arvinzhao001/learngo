package basic

import (
	"fmt"
	"testing"
)

func TestFindUsage(t *testing.T) {
	//The way variadic functions work is by converting the variable number of arguments
	//to a slice of the type of the variadic parameter. for example, convert 1,4,2,3 to
	//[]int{1,4,2,3}, then be passed to the find function
	find(3, 1, 4, 2, 3)

	find(5, 1, 4, 2, 3)

	//nums will be a nil slice with length and capacity 0
	find(3)

	//suffix the slice with ellipsis..., the slice is directly passed to the function
	//without a new slice being created
	nums := []int{89, 90, 95}
	find(90, nums...)
}

func TestChangeInVariadicFunc(t *testing.T) {
	message := []string{"Hello", "World"}
	change(message...)
	fmt.Println("Out function, slice message:", message, " len:", len(message), " cap:", cap(message))
}

//variadic function
func find(num int, nums ...int) {
	fmt.Printf("type of nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Printf("\n")
}

//variadic function, change slice
func change(s ...string) {
	s[0] = "Go"
	//reallocate a new slice
	s = append(s, "Playground")
	fmt.Println("In function, slice message:", s, " len:", len(s), " cap:", cap(s))
}

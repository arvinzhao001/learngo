package basic

import (
	"fmt"
	"testing"
)

func TestCreateMap(t *testing.T) {
	//maps are reference types, map zero value : nil.
	//var employeeSalary map[string]int
	//employeeSalary["arvinzhao"] = 100  //error! employeeSalary is nil, not initialization

	//method 1: use make function
	employeeSalary2 := make(map[string]int)
	employeeSalary2["arvinzhao"] = 100
	employeeSalary2["nathalie"] = 1000
	fmt.Println(employeeSalary2)

	//method 2: initialize a map during the declaration itself
	employeeSalary3 := map[string]int{
		"arvinzhao": 100,
		"nathalie":  1000,
	}
	employeeSalary3["huahua"] = 900
	fmt.Println(employeeSalary3)
}

func TestRetriveMap(t *testing.T) {
	employeeSalary := make(map[string]int)
	employeeSalary["arvinzhao"] = 100
	employeeSalary["nathalie"] = 1000
	employeeSalary["huahua"] = 100
	fmt.Println(employeeSalary)

	fmt.Println("arvinzhao's salary is ", employeeSalary["arvinzhao"])
	/*The map will return the zero value of the type of that element, when the element is not present
	 */
	fmt.Println("sanzhang's salary is ", employeeSalary["sanzhang"])

	// if exsist
	value, ok := employeeSalary["nathalie"]
	if ok == true {
		fmt.Println("nathalie's salary is ", value)
	}

	/*
		The order of the retrieval of values from a map when using "for range" is not guaranteed to be the
		same for each execution of the program.
	*/
	for key, value := range employeeSalary {
		fmt.Printf("employeeSalary[%s]=%d\n", key, value)
	}

	//delete element
	delete(employeeSalary, "huahua")
	fmt.Println(employeeSalary)

	//length of map
	fmt.Println("the number element of employeeSalary: ", len(employeeSalary))

	//maps are reference type
	employeeSalary2 := employeeSalary
	employeeSalary2["arvinzhao"] = 2000
	fmt.Println(employeeSalary)

}

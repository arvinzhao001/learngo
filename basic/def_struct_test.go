package basic

import (
	"fmt"
	"testing"
)

type Address struct {
	province string
	city     string
}

type Grade struct {
	degree string
}

type Employee struct {
	firstName string
	lastName  string
	age       int
	salary    int
	Address
	grade Grade
}

func TestStructBase(t *testing.T) {
	emp6 := Employee{
		firstName: "Sam",
		lastName:  "Anderson",
		age:       55,
		salary:    6000,
		//Promoted field
		Address: Address{
			province: "guangdong",
			city:     "shenzhen",
		},
		//Nested struct
		grade: Grade{
			degree: "master",
		},
	}

	fmt.Println("First Name:", emp6.firstName)
	fmt.Println("Last Name:", emp6.lastName)
	fmt.Println("Age:", emp6.age)
	fmt.Printf("Salary: $%d\n", emp6.salary)
	emp6.salary = 6500
	fmt.Printf("New Salary: $%d\n", emp6.salary)
	fmt.Printf("Province:%s\n", emp6.province)
	fmt.Printf("City:%s\n", emp6.city)
	fmt.Printf("degree:%s\n", emp6.grade.degree)

	emp7 := emp6
	emp7.salary = 7000
	fmt.Println("emp6 salary:", emp6.salary, " emp7 salary:", emp7.salary)

	emp8 := &emp6
	emp8.salary = 7500 //emp8.salary install of the explicit dereference (*emp8).salary to access salary field
	fmt.Println("emp6 salary:", emp6.salary, " emp8 salary:", emp8.salary)
}

func TestZeroValueOfStruct(t *testing.T) {
	var emp Employee
	//Error: can not convert nil to Employee struct
	//fmt.Println(emp == nil)

	//the field of the struct are assigned their zero values by default
	fmt.Println("firstName:", emp.firstName, " age:", emp.age, " province:", emp.province, " degree:", emp.grade.degree)
}

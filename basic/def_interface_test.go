package basic

import (
	"fmt"
	"testing"
)

type Worker interface {
	Work()
}

type Person struct {
	name string
}

func (p Person) Work() {
	fmt.Println(p.name, " is working")
}

//Go interfaces is implemented implicitly if a type contains
//all the methods declared in the interface
func TestInterfaceBase(t *testing.T) {
	p := Person{
		name: "arvinzhao",
	}

	var worker Worker
	//zero value of worker is nil
	if worker == nil {
		worker = p
	}
	worker.Work()
}

//useful example, use interface to improve app's extendibility
type SalaryCalculator interface {
	CalculateSalary() int
}
type Permanent struct {
	empId    int
	basicpay int
	pf       int
}

type Contract struct {
	empId    int
	basicpay int
}

func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.pf
}

func (c Contract) CalculateSalary() int {
	return c.basicpay
}

func totalExpense(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("Total Expense Per Month $%d\n", expense)
}

func TestUseInterface(t *testing.T) {
	pemp1 := Permanent{
		empId:    1,
		basicpay: 5000,
		pf:       20,
	}
	pemp2 := Permanent{
		empId:    2,
		basicpay: 6000,
		pf:       30,
	}
	cemp1 := Contract{
		empId:    3,
		basicpay: 3000,
	}
	cals := []SalaryCalculator{
		pemp1, pemp2, cemp1,
	}
	totalExpense(cals)
}

func TestUnderlyingValueofInterface(t *testing.T) {
	p := Person{
		name: "arvinzhao",
	}
	var worker Worker = p
	worker.Work()

	describe(worker)

	str := "arvinzhao"
	describe(str)

	num := 10
	describe(num)
}

//Since the empty interface has zero methods,
//all types implement the empty interface
func describe(w interface{}) {
	fmt.Printf("Interface type %T value %v\n", w, w)
}

func TestTypeAssertion(t *testing.T) {
	//i.(T) is a syntax which is used to get the underlying value of interface i whose concrete type is T
	num := 10
	var i interface{} = num
	e1 := i.(int)
	fmt.Println("The underlying value of i:", e1)

	e2, ok := i.(string)
	if !ok {
		fmt.Printf("Interface type %T\n", e2)
	}

	findType(&Person{name: "arvinzhao"})
}

/*
type switch
The syntax for type switch is similar to Type assertion. In the syntax i.(T) for Type assertion,
the type T should be replaced by the keyword 'type' for type switch
*/
func findType(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Printf("I am a string and my value is %s\n", i.(string))
	case int:
		fmt.Printf("I am an int and my value is %d\n", i.(int))
	case Worker:
		fmt.Printf("I am an Woker ")
		v.Work()
	default:
		fmt.Printf("Unknown type\n")
	}
}

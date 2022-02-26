package basic

import (
	"fmt"
	"testing"
)

type SalaryCalculator2 interface {
	DisplaySalary()
}

type LeaveCalculator2 interface {
	CalculateLeavesLeft() int
}

type Employee2 struct {
	firstName   string
	lastName    string
	basicPay    int
	pf          int
	totalLeaves int
	leavesTaken int
}

func (e Employee2) DisplaySalary() {
	fmt.Printf("%s %s has salary $%d", e.firstName, e.lastName, (e.basicPay + e.pf))
}

func (e Employee2) CalculateLeavesLeft() int {
	return e.totalLeaves - e.leavesTaken
}

func TestImpMultiInterfaces(t *testing.T) {
	e := Employee2{
		firstName:   "Naveen",
		lastName:    "Ramanathan",
		basicPay:    5000,
		pf:          200,
		totalLeaves: 30,
		leavesTaken: 5,
	}
	var s SalaryCalculator2 = e
	s.DisplaySalary()
	var l LeaveCalculator2 = e
	fmt.Println("\nLeaves left =", l.CalculateLeavesLeft())
}

//create a new interfaces by embedding other interfaces, to simulate inheritance
type EmployeeOperations interface {
	SalaryCalculator2
	LeaveCalculator2
}

func TestSimulateInheritance(t *testing.T) {
	e := Employee2{
		firstName:   "Naveen",
		lastName:    "Ramanathan",
		basicPay:    5000,
		pf:          200,
		totalLeaves: 30,
		leavesTaken: 5,
	}

	var empOp EmployeeOperations = e
	empOp.DisplaySalary()
	fmt.Println("\nLeaves left =", empOp.CalculateLeavesLeft())
}

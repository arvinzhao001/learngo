package basic

import (
	"fmt"
	"testing"
)

//unimport type employee, force caller uses New Method to initailize object
type employee struct {
	firstName   string
	lastName    string
	totalLeaves int
	leavesTaken int
}

//import mock constructor
func New(firstName string, lastName string, totalLeave int, leavesTaken int) employee {
	e := employee{firstName, lastName, totalLeave, leavesTaken}
	return e
}

func (e employee) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining\n", e.firstName, e.lastName, (e.totalLeaves - e.leavesTaken))
}

func TestMockOpp(t *testing.T) {
	e := New("arvin", "zhao", 30, 20)
	e.LeavesRemaining()
}

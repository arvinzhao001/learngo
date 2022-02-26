package basic

import (
	"fmt"
	"testing"
)

type Describer interface {
	Describe()
}

type Person2 struct {
	name string
	age  int
}

//implemented using value receiver
func (p Person2) Describe() {
	fmt.Printf("%s is %d years old\n", p.name, p.age)
}

type Address2 struct {
	state   string
	country string
}

//implemented using pointer receiver
func (a *Address2) Describe() {
	fmt.Printf("State %s Country %s", a.state, a.country)
}

func TestDiffReceiverForInterface(t *testing.T) {
	var d1 Describer
	p1 := Person2{name: "arvinzhao", age: 32}
	d1 = p1
	d1.Describe()
	d1 = &p1
	d1.Describe()

	a1 := Address2{state: "Guangdong", country: "Shenzhen"}

	//The reason is that it is legal to call a pointer-valued method
	//on anything that is already a pointer or whose address can be taken.
	//The concrete value stored in an interface is not addressable and
	//hence it is not possible for the compiler to automatically take the address of a1 and hence this code fails.
	//d1 = a1
	d1 = &a1
	d1.Describe()
}

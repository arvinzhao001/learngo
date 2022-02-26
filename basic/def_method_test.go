package basic

import (
	"fmt"
	"math/rand"
	"testing"
)

//to define a method on a type, the definition of the reciever type and
//the definition of the method shold be present in the same package
type Test struct {
	value int
}

//copy a data structure, changes is not visible to the caller
func (test Test) changeValue1() {
	test.value = 10
}

//changes is visible to the caller
func (test *Test) changeValue2() {
	test.value = 20 + rand.Intn(100)
}

func TestValueVsPointer(t *testing.T) {
	var test = Test{
		value: 0,
	}

	fmt.Println("test:", test.value)
	//value reciever
	test.changeValue1()
	fmt.Println("after call value reciever, test:", test.value)
	//for convenience will be interpreted by GO as (&test).changeValue2
	test.changeValue2()
	fmt.Println("after call pointer reciever, test:", test.value)

	var test2 *Test = &test
	//for convenience will be interpreted by GO as (*test2).changeValue1
	test2.changeValue1()
	fmt.Println("after call value reciever, test:", test.value)
	//pointer reciever
	test2.changeValue2()
	fmt.Println("after call pointer reciever, test:", test.value)

}

type Test2 struct {
	Test
}

func TestMethodOfAnonStructField(t *testing.T) {
	test2 := Test2{
		Test: Test{
			value: 0,
		},
	}
	(&test2).changeValue2()
	fmt.Println("after call method of anonymous field:", test2.value)
}

package basic

import (
	"fmt"
	"learngo/interest"
	"log"
	"testing"
)

var p, r, t = 5000.0, 10.0, 1.0

func init() {
	println("Self package initialized")
	if p < 0 {
		log.Fatal("Principal is less than zero")
	}
	if r < 0 {
		log.Fatal("Rate of interest is less than zero")
	}
	if t < 0 {
		log.Fatal("Duration is less than zero")
	}
}

/*
1. package level variables are initialized first,init function is called next.
2. a package can have multiple init functions, they are called in the order which they are prensented to the compiler
3. if a package imports other packeges, the imported package are initialized first
4. a package will be initialized only once even if it is imported from multiple packages
*/
func TestPackageInitialization(c *testing.T) {
	fmt.Println("Simple interest calculation")
	si := interest.Calculate(p, r, t)
	fmt.Println("Simple interest is", si)
}

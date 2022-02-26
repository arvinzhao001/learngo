package main

//tutorials : https://golangbot.com/concurrency/
import (
	"fmt"

	"github.com/arvinzhao001/learngo/basic"
	"github.com/arvinzhao001/learngo/lib/copier"
)

func init() {
	fmt.Println("Main package is initalized")
}

func main() {
	copier.TestCopier()
	basic.TestScrypt()
}

package main

//tutorials : https://golangbot.com/concurrency/
//Go module : https://roberto.selbach.ca/intro-to-go-modules/
import (
	"fmt"

	"github.com/arvinzhao001/learngo/lib/copier"
	"github.com/arvinzhao001/learngo/lib/secret"
)

func init() {
	fmt.Println("Main package is initalized")
}

func main() {
	copier.TestCopier()
	secret.TestScrypt()
}

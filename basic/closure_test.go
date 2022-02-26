package basic

import (
	"fmt"
	"testing"
)

//high order function
func imap(datas []int, f func(d int) int) []int {
	var result []int
	for _, data := range datas {
		result = append(result, f(data))
	}
	return result
}

func TestClosure(t *testing.T) {
	datas := []int{1, 2, 3, 4, 5}
	result := imap(datas, func(d int) int {
		return d * d
	})
	fmt.Println(result)
}

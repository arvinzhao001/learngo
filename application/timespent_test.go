package basic

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

type target func(int, int) int

func TimeSpent(f target) target {
	return func(start, end int) int {
		defer func(t time.Time) int {
			fmt.Printf("Function: %s, Spend Time: %v\n", GetFunctionName(f), time.Since(t))
			return f(start, end)
		}(time.Now())
		return f(start, end)
	}
}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func Sum1(start, end int) int {
	sum := 0
	if start > end {
		start, end = end, start
	}

	for i := start; i <= end; i++ {
		sum += i
	}

	return sum
}

func Sum2(start, end int) int {
	if start > end {
		start, end = end, start
	}
	return (end - start + 1) * (end + start) / 2
}

func TestTimeSpend() {
	f1 := TimeSpent(Sum1)
	f2 := TimeSpent(Sum2)

	sum1 := f1(100, 10000)
	sum2 := f2(100, 10000)
	fmt.Printf("Sum1: %d, Sum2: %d", sum1, sum2)
}

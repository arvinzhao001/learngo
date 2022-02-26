package basic

import (
	"fmt"
	"testing"
)

type Season int8

const (
	Season_DEFAULT Season = iota
	Season_SPRING
	Season_SUMMER
	Season_AUTOMN
	Season_WINTER
)

func (s Season) String() string {
	switch s {
	case Season_SPRING:
		return "spring"
	case Season_AUTOMN:
		return "automn"
	case Season_SUMMER:
		return "summer"
	case Season_WINTER:
		return "winter"
	default:
		return "undefined"
	}
}

func TestEnum(t *testing.T) {
	fmt.Println("Season_DEFAULT code: ", int8(Season_DEFAULT))
	fmt.Println("Season_DEFAULT flag: ", Season_DEFAULT.String())

	fmt.Println("Season_WINTER code: ", int8(Season_WINTER))
	fmt.Println("Season_WINTER flag: ", Season_WINTER.String())
}

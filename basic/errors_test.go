package basic

import (
	"errors"
	"fmt"
	"net"
	"testing"
)

//the built in error type is a interface type with the following definition.
//contains a single method with signature Error(). Any type which implements
//this interface can be used as a error
/*
type error interface {
    Error() string
}
*/
func TestErrorBase(t *testing.T) {
	addr, err := net.LookupHost("golangbot123.com")
	if err != nil {
		//we use type assertion to get the underlying value of the error interface
		if dnsErr, ok := err.(*net.DNSError); ok {
			if dnsErr.Timeout() {
				fmt.Println("operation timed out")
				return
			}
			if dnsErr.Temporary() {
				fmt.Println("temporary error")
				return
			}
			fmt.Println("Generic DNS error", err)
			return
		}
		fmt.Println("Generic error", err)
		return
	}
	fmt.Println(addr)
}

//create customed error use built-in lib
func createErr1() error {
	return errors.New("test customed error")
}

func createErr2() error {
	hello := "hello"
	return fmt.Errorf("test customed error:%s", hello)
}

func TestCustomizedError1(t *testing.T) {
	err1 := createErr1()
	fmt.Println(err1)
	err2 := createErr2()
	fmt.Println(err2)
}

//customized error type
type areaError struct {
	err    string //error description
	length float64
	width  float64
}

func (e *areaError) Error() string {
	return e.err
}

func (e *areaError) lengthNegative() bool {
	return e.length < 0
}

func (e *areaError) widthNegative() bool {
	return e.width < 0
}

func rectArea(length, width float64) (float64, error) {
	err := ""
	if length < 0 {
		err = "length is less than zero"
	}
	if width < 0 {
		if err == "" {
			err = "width is less than zero"
		} else {
			err += ", width is less than zero"
		}
	}
	if err != "" {
		return 0, &areaError{err, length, width}
	}
	return length * width, nil
}

func TestCustomizedError2(t *testing.T) {
	length, width := -10.0, -20.0
	area, err := rectArea(length, width)
	if err != nil {
		if err, ok := err.(*areaError); ok {
			if err.lengthNegative() {
				fmt.Printf("error: length %0.2f is less than zero\n", err.length)
			}
			if err.widthNegative() {
				fmt.Printf("error: width %0.2f is less than zero\n", err.width)
			}
		}
		return
	}

	fmt.Println("Area is ", area)
}

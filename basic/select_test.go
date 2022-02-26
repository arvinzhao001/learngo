package basic

import (
	"fmt"
	"testing"
	"time"
)

func server1(ch chan string) {
	time.Sleep(1500 * time.Millisecond)
	ch <- "from server1"
}
func server2(ch chan string) {
	time.Sleep(1500 * time.Millisecond)
	ch <- "from server2"
}
func server3(ch chan string) {
	time.Sleep(3000 * time.Millisecond)
	ch <- "from server3"
}

func TestSelectBase(t *testing.T) {
	s1 := make(chan string)
	s2 := make(chan string)
	s3 := make(chan string)
	go server1(s1)
	go server2(s2)
	go server3(s3)
	for {
		time.Sleep(1000 * time.Millisecond)
		/*
			1. The select statement is used to choose from multiple send/recieve channel operations,
			blocks until one of the send/recieve operations is ready
			2. The default case in a select statement is executed when none of the other cases is ready
			3. When multiple cases in a select statement are ready, one of them will be executed at random
		*/
		select {
		case v1 := <-s1:
			fmt.Println("received value: ", v1)
			return
		case v2 := <-s2:
			fmt.Println("received value: ", v2)
			return
		case v3 := <-s3:
			fmt.Println("received value: ", v3)
			return
		default:
			fmt.Println("no value received")
		}
	}

}

func closeChannel(ch chan string) {
	time.Sleep(1500 * time.Millisecond)
	close(ch)
}
func TestCloseChannel(t *testing.T) {
	s1 := make(chan string)
	go closeChannel(s1)
	for {
		select {
		case <-s1:
			fmt.Println("channel closed")
		default:
			fmt.Println("channel active")
		}
		time.Sleep(1000 * time.Millisecond)
	}

}

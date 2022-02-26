package basic

import (
	"fmt"
	"testing"
	"time"
)

func TestChannelBase(t *testing.T) {
	//chan T is a channel of type T
	var a chan int
	//zero value of channel is nil
	if a == nil {
		a = make(chan int)
	}
	fmt.Printf("Type of a is %T\n", a)

	// sends and receives are blocking by default
	// when data is sent to a channel, the control is blocked in the send
	// statement until some other Goroutine reads from that channel. Read
	// operation is similar
	// reader := <-a
	// a <- writer
}

func TestChannelSync(t *testing.T) {
	number := 589
	sqrch := make(chan int)
	cubch := make(chan int)
	go calSquares(number, sqrch)
	go calCubes(number, cubch)

	squares, cubes := <-sqrch, <-cubch
	fmt.Println("Final output: ", squares+cubes)
}

func TestConsumerProducer(t *testing.T) {
	chnl := make(chan int)
	go producer(chnl)
	//Two consumer consume chnl data currently
	go consumer1(chnl)
	go consumer2(chnl)

	time.Sleep(2 * time.Second)
}

//unidirectional channel: squareop
//It is possible to convert a bidirectional channel to a send
//only or receive only channel but not the vice versa
func calSquares(number int, squareop chan<- int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum = sum + digit*digit
		number = number / 10
	}
	squareop <- sum
}

func calCubes(number int, cubeop chan<- int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum = sum + digit*digit*digit
		number = number / 10
	}
	cubeop <- sum
}

func producer(chnl chan<- int) {
	for i := 1; i < 10; i++ {
		chnl <- i
		fmt.Println("Producer send ", i)
	}
	//close channel
	//Senders have the ability to close the channel to notify receivers that no more data will
	//be sent on the channel
	close(chnl)
}

//judge if the channel has closed, method 1
func consumer1(chnl chan int) {
	for {
		v, ok := <-chnl
		if ok == false {
			break
		}
		fmt.Println("Comsumer1 recieved ", v)
		time.Sleep(10 * time.Millisecond)
	}
}

//judge if the channel has closed, method 2
func consumer2(chnl chan int) {
	for v := range chnl {
		fmt.Println("Consumer2 recieved ", v)
		time.Sleep(10 * time.Millisecond)
	}
}

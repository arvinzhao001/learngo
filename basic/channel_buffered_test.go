package basic

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

type Job struct {
	id       int
	randomno int
}
type Result struct {
	job         Job
	sumofdigits int
}

//Sends to a buffered channel are blocked only when the buffer is full.
//Similarly receives from a buffered channel are blocked only when the buffer is empty
var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

//The capacity of a buffered channel is the number of values that the channel can hold
//The length of a buffered channel is the number of elements currently queued in it
func TestBufferedChannelCap(t *testing.T) {
	fmt.Println("Capicity is ", cap(jobs), ", Length is ", len(jobs))
}

func TestWorkerPool(t *testing.T) {
	startTime := time.Now()
	noOfJobs := 100
	go allocate(noOfJobs)
	//the capacity for an unbuffered channel is 0 by default
	done := make(chan bool)
	go result(done)
	noOfWorkers := 10
	createWorkerPool(noOfWorkers)
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}

func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(2 * time.Second)
	return sum
}

func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, digits(job.randomno)}
		results <- output
	}
	wg.Done()
}

//A WaitGroup is used to wait for a collection of Goroutines to finish executing
func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	//It's possible to read data from a closed buffered channel.
	//The channel will return the data that is already written to the channel
	//and once all the data has been read, it will return the zero value of the channel
	close(results)
}
func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		jobs <- job
	}
	close(jobs)
}
func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
	}
	done <- true
}

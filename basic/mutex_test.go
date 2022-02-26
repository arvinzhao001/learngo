package basic

import (
	"fmt"
	"sync"
	"testing"
)

var x = 0
var y = 0

//In general use channels when Goroutines need to communicate with each
//other and mutexes when only one Goroutine should access the critical section of code
func TestSolveRaceByMutex(t *testing.T) {
	var w sync.WaitGroup
	var m sync.Mutex

	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w, &m)
	}

	w.Wait()
	fmt.Println("final value is ", x)
}

func TestSolveRaceByChan(t *testing.T) {
	var w sync.WaitGroup
	//one sized buffered channel
	ch := make(chan bool, 1)

	for i := 0; i < 500; i++ {
		w.Add(1)
		go increment2(&w, ch)
	}

	w.Wait()
	fmt.Println("final value is ", y)
}

//solve race condition by mutex
func increment(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x = x + 1
	m.Unlock()
	wg.Done()
}

//solve race condition by channel
func increment2(wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	y = y + 1
	<-ch
	wg.Done()
}

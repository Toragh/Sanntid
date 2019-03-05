// Use `go run foo.go` to run your program

package main

import (
	. "fmt"
	"runtime"
)

// States
type State int;
const (
	GetNumber = iota
	Exit
)

func number_server(add_number <-chan int, control <-chan int, number chan<- int) {
	var i = 0

	// This for-select pattern is one you will become familiar with if you're using go "correctly".
	for {
		select {
		// TODO: receive different messages and handle them correctly
		// You will at least need to update the number and handle control signals.
		case n := <-add_number:
			i += n
		case msg := <-control:
			if msg == GetNumber {
				number <- i
			} else if msg == Exit {
				return
			}
		}
	}
}

func incrementing(add_number chan<- int, finished chan<- bool) {
	for j := 0; j < 1000000; j++ {
		add_number <- 1
	}
	//TODO: signal that the goroutine is finished
	finished <- true
}

func decrementing(add_number chan<- int, finished chan<- bool) {
	for j := 0; j < 1000000; j++ {
		add_number <- -1
	}
	//TODO: signal that the goroutine is finished
	finished <- true
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	finished := make(chan bool)
	control := make(chan int)
	number := make(chan int)
	add_number := make(chan int)

	go decrementing(add_number, finished)
	go incrementing(add_number, finished)
	go number_server(add_number, control, number)
	<-finished
	<-finished

	control <- GetNumber
	Println("The magic number is:", <-number)
	control <- Exit
}

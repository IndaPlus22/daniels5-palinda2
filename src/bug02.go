package main

import (
	"fmt"
	"sync"
	"time"
)

//Problem: Is that the main function returns before the print function manages to print all the values. When the main routine returns it kills all the other goroutines this means that the print routine won't have enough time to print all the intigers
//Solution: We solve this issue using sync.WaitGroup. First we tell the main goroutine to add 1 routine too wait for, then att the end before it returns we tell the main routine too wait for the other routine to be done.
//In the print routine we defer a wg.Done() so that it calls that this routine is done at the end after that the main routine can safely return.

// This program should go to 11, but it seemingly only prints 1 to 10.
var wg sync.WaitGroup

func main() {
	wg.Add(1)
	ch := make(chan int)
	go Print(ch)
	for i := 1; i <= 11; i++ {
		ch <- i
	}
	close(ch)
	wg.Wait() //Wait for the print routine

}

// Print prints all numbers sent on the channel.
// The function returns when the channel is closed.
func Print(ch <-chan int) {
	defer wg.Done()     //This makes the print function call wg.Done once everything else has ran in the function.
	for n := range ch { // reads from channel until it's closed
		time.Sleep(10 * time.Millisecond) // simulate processing time
		fmt.Println(n)
	}
}

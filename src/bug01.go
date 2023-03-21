package main

import "fmt"

//PROBLEM: Channel receiving and sending acts as a synchronization point between different go routines. For example a channel recieve will wait until another go routine is sending to channel, and the go routine sending is waiting for a go routine to recieve.
//Here both the sending and recieving is happening on the same routine or in other words sequentily, this causes the send to wait forever for a recieve since. Hence a deadlock

//SOLUTION: Either we could directly print hello world or the solution i'am using seperate the channel send into another function and then send that to a go routine, that way we can synchronize the 2 routines properly.

func main() {
	ch := make(chan string)

	go func() {
		ch <- "Hello world!"
	}()

	fmt.Println(<-ch)
}

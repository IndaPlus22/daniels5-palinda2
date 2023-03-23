If written answers are required, you can add them to this file. Just copy the
relevant questions from the root of the repo, preferably in
[Markdown](https://guides.github.com/features/mastering-markdown/) format :)

#### Task 1

##### Buggy Code 1
1. PROBLEM: Channel receiving and sending acts as a synchronization point between different go routines. For example a channel recieve will wait until another go routine is sending to channel, and the go routine sending is waiting for a go routine to recieve.
Here both the sending and recieving is happening on the same routine or in other words sequentily, this causes the send to wait forever for a recieve since. Hence a deadlock

2. SOLUTION: Either we could directly print hello world or the solution i'am using seperate the channel send into another function and then send that to a go routine, that way we can synchronize the 2 routines properly.

##### Buggy Code 2
1. Problem: Is that the main function returns before the print function manages to print all the values. When the main routine returns it kills all the other goroutines this means that the print routine won't have enough time to print all the intigers

2. Solution: We solve this issue using sync.WaitGroup. First we tell the main goroutine to add 1 routine too wait for, then att the end before it returns we tell the main routine too wait for the other routine to be done.
In the print routine we defer a wg.Done() so that it calls that this routine is done at the end after that the main routine can safely return.

// This program should go to 11, but it seemingly only prints 1 to 10.

#### Task 2

| Question | What I expected | What happened | Why I believe this happened |
|-|-|-|-|
| What happens if you do X? |  Program would still work as before | Program ended up in a deadlock | Because of reasons 🤷 |
| What happens if you switch the order of the statements `wgp.Wait()` and `close(ch)` in the end of the `main` function? | I expect it to to error out | A error happened that states that a goroutine is trying to send to a closed channel | This happens because the main routine closes the channel before the produce routines manages to send the relevant data to the channel, in other words the produce routines don't have enough time to send the data before it closes |
| What happens if you move the `close(ch)` from the `main` function and instead close the channel in the end of the function `Produce`?  | I expect it to manage to send a few strings and then crash | That is exactly what happend | Since the close function is now in each produce routine then the first produce routine too finnish will close the channel that all routines use, which might cause the remaining 3 routines to try to send to a closed channel. |
| What happens if you remove the statement `close(ch)` completely?  | I expect a deadlock since the recive routine will trying to recieve forever | It works the same way as if we still had the close function | This might be because the for loop used by the reciever routines might be because the reciever routines are killed when the main routine exits thus preventing them from waiting forever |
| What happens if you increase the number of consumers from 2 to 4?  | I expect it to work as usuall | This was the case | Since the consumers will still be sharing the workload among themselfs when recieving the produced strings |
| Can you be sure that all strings are printed before the program stops?  | No | ... | This is because the consumers will print the strings aslong as the channel is open AND that the main go routine is still active, However the main go routine might exit before the consumers have enough time to print all the strings |

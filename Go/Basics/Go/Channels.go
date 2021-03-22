package main

import (
	"fmt"
	"time"
	"sync"
)

// WAIT GROUPS
// the main method does not know to wait for go routines by default
// it will simply enter the main method, run through to the end and terminate
// providing it's run any normal functions (i.e. not go routines)
// as such, you can use wait groups to tell the main method to wait until all threads etc. have run
// here we define the wait group object which we will invoke in the main method itself
// the steps are as follows
// 1) create wait group variable (this line here)
// 2) add a wait group to the start of the main method
// 3) add a wait group .Done() method to the relevant function/go routine where the work is taking place
// 4) use the .Wait() method at the end of the main method to check that everything has run before terminating
var wg sync.WaitGroup

// SYNCHRONIZATION
// let's define a function that simply prints 2 strings a few seconds apart
// we add a line at the bottom which sends a signal to the channel stating that it's done
func worker(done chan bool) {
	defer wg.Done() // wait until below code has run, then signal to the wait group that work is complete
	fmt.Println("Starting...") // first message
	time.Sleep(3 * time.Second) // simulate work
	fmt.Println("...finished") // last message

	done <- true // send boolean signal to channel (input to this function)
}

// CHANNEL DIRECTIONS
// you can specify whether a channel can send only, receive only or send and receive
// this is done using functions with input vars equal to channels with specific criteria
// this function is designed to send a string to a channel but receive nothing
// the chan<- syntax means "channel send"
func sendOnly(c1 chan<- string, msg string) {
	// here, we send the input message to the input channel
	c1 <- msg
}

// this function can both send and receive a string
// again, the chan<- means it can send a string, whilst the <-chan means it can receive a string
func sendAndReceive(rec <-chan string, sen chan<- string) {
	// receive message from channel which is able to receive (i.e. <-chan)
	msg := <-rec

	// send message to channel which is able to send (i.e. chan<-)
	sen <- msg
}

func main() {
	// WAIT GROUP
	// add wait group at start of main method to invoke it
	// having a positive delta means it will wait until all go routines run
	wg.Add(1)

	// NOTES
	// go routines are lightweight threads i.e. they allow multiple concurrent processes to run
	// channels are specific objects which allow you to control specific processes concurrently with others
	// the below code covers how to create, manage and control go routines and channels

	// UNBUFFERED CHANNELS (DEFAULT)
	// create channel called messages which expects strings
	// channels are simply ways of specifying exact pathways for
	// variables/routines etc. to pass along, allowing thread control
	messages := make(chan string)

	// send value to channel via go routine anonymous function
	// <- is the send operator
	go func() {messages <- "Test"}()

	// receive value into var and show
	// := <- defines the receive operator(s)
	msg := <-messages
	fmt.Println(msg)

	// BUFFERED CHANNELS
	// the above is an unbuffered channel, which is the default
	// this means that for every send signal, there must be a receive
	// you can create unbuffered channels which do not require a receive
	messages2 := make(chan string, 2) // define a channel expecting exactly 2 values/signals
	messages2 <- "Unbuffered" // send signal 1
	messages2 <- "Channel" // send signal 2
	fmt.Println(<-messages2) // able to retrieve both signals without requiring a receive method
	fmt.Println(<-messages2)

	// SYNCHRONIZATION
	// you can build channels so that they wait for a signal of completion
	// this prevents channels closing or not running before the work is done
	done := make(chan bool) // create channel expecting 1 bool
	go worker(done) // create go routine, passing the worker function the above channel

	// wait for true bool signal from worker function
	// if this line wasn't here, the worker wouldn't run
	// the code would just stop as it's not expecting anything
	<- done


	// CHANNEL DIRECTIONS
	c1 := make(chan string, 1) // create channel 1, expecting 1 string
	c2 := make(chan string, 1) // create channel 2, expecting 1 string

	// below lines essentially are doing the following
	// msg > send to c1 > get msg from c1 > send msg to c2 > get msg from c2 > print
	sendOnly(c1, "Send this") // send message via channel 1
	sendAndReceive(c1, c2) // receive the message from c1 (sent in above line) and receive it to channel 2
	fmt.Println(<-c2) // extract received message from channel 2

	// SELECT
	// you can use to select to control the results for different channels
	// firstly we'll create the 2 channels themselves
	ca := make(chan string)
	cb := make(chan string)

	// now we'll create an anonymous function which waits for 1 second before returning a result
	go func() {
		time.Sleep(1 * time.Second) // wait for 1 second
		ca <- "one"
	}()

	// we'll create another identical function except that it will wait for 5 seconds
	go func() {
		time.Sleep(10 * time.Second) // wait for 5 seconds
		cb <- "two"
	}()

	// we can then use select to show the messages as they're received
	// note that the second message will come back after a few seconds
	// we iterate twice here to get the 2 messages we're expecting
	// we only receive the first message, because our timeout case checks
	// for responses after 2 seconds and if it's not received one it is satisfied
	// because our cb result comes after 10 seconds, the timeout is triggered before it arrives
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ca:
			fmt.Println("Received: ", msg1)
		case msg2 := <-cb:
			fmt.Println("Received: ", msg2)
		// <-time.After() means it triggers after the time specified
		// here, this case is satisfied after 2 seconds
		case <-time.After(2 * time.Second):
			fmt.Println("Timeout error")
		// we can use default to handle cases where the above cases aren't achieved
		// commented out here because default will always trigger as select won't
		// wait for the timed delay responses, it will just run
		//default:
		//	fmt.Println("No message received")
		}
	}

	// CLOSING CHANNELS AND RANGE ITERATION
	// once we're sure a channel is complete, we can close it for safety
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// we can iterate over channel values using range
	// this is possible even after the channel is closed
	for elem := range queue {
		fmt.Println(elem)
	}

	// TIMERS
	// timers can be used to wait until a specific time in the future
	// this is useful if you're looking to trigger something once at a specific time
	// below we create 2 timer objects, setting them to 2 and 3 seconds in the future respectively
	timer := time.NewTimer(2 * time.Second)
	timer2 := time.NewTimer(3 * time.Second)

	// we can then trigger an action after the timer has run
	// the below line (<-timer.C) simply means block the timer's channel (c) until the specified duration (i.e. 2 secs)
	<-timer.C
	fmt.Println("Timer fired")

	// we could simply use time.Sleep(duration) to achieve the same result
	// but timers let us stop them before they fully run if certain conditions are met too
	// this function will attempt to wait for timer2 to wait for the specified 3 seconds
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()

	// however, this step will force the timer to stop before it's reached 3 seconds
	// because the above function is a go routine so will run concurrently with the below
	// as such, the below message will fire and the above won't as the timer won't complete its duration
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	// this line just ensures that there is plenty of time for the above to run to completion
	time.Sleep(5 * time.Second)

	// TICKERS
	// ticker are like timers except that they allow you to do something multiple times in the future
	// firstly we'll create a ticker that does something every 500ms and a channel that expects a bool
	ticker := time.NewTicker(500 * time.Millisecond)
	done2 := make(chan bool)

	// now let's create an anonymous function which prints a time until a true bool is sent to the channel
	go func() {
		for {
			select {
				// if channel returns a true bool
				case <-done2:
					// finish
					return
				// otherwise get the value from my ticker (i.e. time)
				case t := <-ticker.C:
					// and print the timestamp
					fmt.Println("Tick at: ", t)
			}
		}
	}()

	// delay everything for 1600 milliseconds, allowing the ticker to fire 3 times (i.e. 500 x 3)
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop() // force the ticker to stop
	done2 <- true // send the true bool to the channel, triggering the go routine above to stop
	fmt.Println("Ticker stopped") // confirm ticker stop

	// WAIT GROUP
	// final element of the wait group, wait here until above wait group stages are complete
	// i.e. create wg > add > done > wait > exit main method
	// note that if you don't have this code, you will receive a fatal error that your go routines are asleep
	wg.Wait()
	fmt.Println("Work complete")
}

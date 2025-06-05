// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"time"
)

func main() {
	//write func that recive a slice of channels int and listen all channels to send first recived value from any channel to output. output is channel int.
	fmt.Println("Hello, 世界")
	//create channels slice - example
	channels := []chan int{
		make(chan int),
		make(chan int),
		make(chan int),
	}

	//send some data to channel #0
	go func() {
		time.Sleep(1 * time.Second)
		channels[0] <- 11
		time.Sleep(3 * time.Second)
		channels[0] <- 14

	}()

	//send some data to channel #1
	go func() {
		time.Sleep(2 * time.Second)
		channels[1] <- 22
		time.Sleep(3 * time.Second)
		channels[1] <- 25

	}()

	//send some data to channel #2
	go func() {
		time.Sleep(3 * time.Second)
		channels[2] <- 33
		time.Sleep(3 * time.Second)
		channels[2] <- 36

	}()

	//send immidiatly first recived number
	//to make slice listen need to rewrite to reflection .
	for range 6 {
		select {
		case msg1 := <-channels[0]:
			fmt.Println("recived", msg1)
		case msg2 := <-channels[1]:
			fmt.Println("recived", msg2)
		case msg3 := <-channels[2]:
			fmt.Println("recived", msg3)
		}
	}
}

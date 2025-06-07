// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"reflect"
	"time"
)

func mass(channels []chan int, out chan int) {

	go func(channels []chan int, out chan int) {
		// max length in 65535 values.
		cases := make([]reflect.SelectCase, len(channels))
		for j := 0; j < len(channels); j++ {
			cases[j] = reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(channels[j]),
				Send: reflect.Value{},
			}
		}
		//we can add more checks to confirm that all channels have been closed. But for this example enough add 6 as number of messages.
		for range 6 {
			// wait until recive one of channels responce
			chIndex, val, ok := reflect.Select(cases)
			if ok {
				fmt.Println("log: chanel arrieved ", chIndex, " value:", val)
				//pause gorutine untill chanel readed
				out <- int(val.Int())

			}
		}
		//close output channel to inform about last message and prevent deadlock
		close(out)

	}(channels, out)

}

func main() {
	fmt.Println("Hello, 世界", " This is an example of test task for lifecoding. recive messges from slice of channels and send immidiatly to output channel.")
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
	out := make(chan int, 1)
  //run function that is a goal of this task
	mass(channels, out)

	for val := range out {
		fmt.Println("Received:", val)
	}

}

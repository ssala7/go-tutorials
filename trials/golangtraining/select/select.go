package main

import "time"
import "fmt"

func main() {

	// For our example we'll select across two channels.
	c1 := make(chan string)
	c2 := make(chan string)

	// Each channel will receive a value after some amount
	// of time, to simulate e.g. blocking RPC operations
	// executing in concurrent goroutines.

	go func() {
		for i := 0; i < 4; i++ {
			time.Sleep(time.Microsecond)
			c1 <- "image" + string(i)
		}
		close(c1)
	}()

	go func() {
		time.Sleep(time.Millisecond)
		c2 <- "video"
		close(c2)
	}()
	// go func() {
	// 	time.Sleep(time.Millisecond)
	// 	c3 <- "text"
	// }()
	// go func() {
	// 	time.Sleep(time.Millisecond)
	// 	c4 <- "audio"
	// }()

	// We'll use `select` to await both of these values
	// simultaneously, printing each one as it arrives.
	for {
		select {
		case msg1, ok := <-c1:
			fmt.Println("received img", msg1)
			if !ok {
				c1 = nil
			}
		case msg2, ok := <-c2:
			fmt.Println("received vdo", msg2)
			if !ok {
				c2 = nil
			}
		}
		if c1 == nil && c2 == nil {
			break
		}
	}
}

package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			c1 <- "Every 500 ms"
		}

	}()

	go func() {

		for {
			time.Sleep(time.Second * 2)
			c2 <- "Every 2s"
		}

	}()

	for {
		select { // select the msg from channel which is ready
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}

}

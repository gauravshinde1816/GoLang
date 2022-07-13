package main

import (
	"fmt"
	"time"
)

func main() {
	go count("Sheep")
	go count("Dog")

	fmt.Scanln()

}

func count(thing string) {
	for i := 1; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500) //sleep for 500 ms
	}

}

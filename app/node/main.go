package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("BlinkNode Agent is starting...")

	fmt.Println("BlinkNode Agent Successfully Started")
	for {
		fmt.Println("Running... ")
		time.Sleep(5 * time.Second)
	}
}

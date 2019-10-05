package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println("say ", s, i)
		time.Sleep(time.Duration(1 * time.Second))
	}
}

func main() {
	fmt.Println("Start")
	go say("Hi") // run concurrency

	// if main ended before concurrent process, all thread will be kill before proceed also
	time.Sleep(time.Duration(7 * time.Second))

	fmt.Println("End")
}

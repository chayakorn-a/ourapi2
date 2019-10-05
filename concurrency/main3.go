package main

import (
	"fmt"
	"sync"
	"time"
)

func say(s string, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		fmt.Println("say ", s, i)
		time.Sleep(time.Duration(1 * time.Second))
	}
}

func main() {
	fmt.Println("Start")

	c := make(chan string)

	go func() {
		fmt.Println("Hello")
		data := <-c
		fmt.Println("data in channel: ", data)
	}()

	fmt.Println("before send")
	time.Sleep(2 * time.Second)
	c <- "Some data."
	time.Sleep(2 * time.Second)
	fmt.Println("Sent")
}

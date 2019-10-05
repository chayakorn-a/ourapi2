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

	var wg sync.WaitGroup
	wg.Add(2)

	go say("Hi", &wg) // run concurrency
	go say("Test", &wg)

	// if main ended before concurrent process, all thread will be kill before proceed also
	//time.Sleep(time.Duration(7 * time.Second))
	wg.Wait()
	fmt.Println("End")
}

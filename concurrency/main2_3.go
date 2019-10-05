package main

import (
	"fmt"
	"time"
)

func say(s string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; i < 5; i++ {
			c <- fmt.Sprintf("say %s: %d", s, i)
			time.Sleep(time.Duration(1 * time.Second))
		}
	}()

	return c
}

func main() {
	c := say("Hi")

	for i := 0; i < 5; i++ {
		fmt.Println("Get word: ", <-c)
	}

	fmt.Println("End")
}

package main

import (
	"fmt"
	"time"
)

func say(s string, c chan string) {
	for i := 0; i < 5; i++ {
		c <- fmt.Sprintf("say %s: %d", s, i)
		time.Sleep(time.Duration(1 * time.Second))
	}
}

func main() {
	c := make(chan string)

	go say("Hi", c)

	for i := 0; i < 5; i++ {
		fmt.Println("You say: ", <-c)
	}

	fmt.Println("End")
}

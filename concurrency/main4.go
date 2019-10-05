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

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)

	/*go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()*/

	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case y := <-input2:
				c <- y
			case <-time.After(3000 * time.Millisecond):
				fmt.Println("too slow")
				close(c)
				return
			}
		}
	}()

	return c
}

func main() {
	c1 := say("Hi")
	c2 := say("Hello")

	c := fanIn(c1, c2)

	for i := 0; i < 10; i++ {
		if data, ok := <-c; ok {
			fmt.Println("Get word: ", data)
		}
	}

	fmt.Println("End")
}

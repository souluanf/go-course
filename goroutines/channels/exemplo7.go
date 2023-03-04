package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Fun in
func main() {
	x := funnel(generateMessage("hello go"), generateMessage("hello world"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-x)
	}
	fmt.Println("Finished")

}

func generateMessage(s string) <-chan string { // ele vai retornar um channel que só receb informação
	channel := make(chan string)
	go func() {
		for i := 0; ; i++ {
			channel <- fmt.Sprintf("String: %s - Value: %d", s, i)
			time.Sleep(time.Duration(rand.Intn(255)) * time.Millisecond)
		}
	}()
	return channel
}

func funnel(channel1, channel2 <-chan string) <-chan string {
	channel := make(chan string)
	go func() {
		for {
			channel <- <-channel1
		}
	}()
	go func() {
		for {
			channel <- <-channel2
		}
	}()

	return channel
}

/**

// P1 -> channel
// P2 -> channel
// P3 -> channel

function P1,P2,P3
return channelX -> P1,P2,P3

*/

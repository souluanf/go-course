package main

import (
	"fmt"
	"time"
)

func main() {
	chanel := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			chanel <- i
		}
	}()
	go func() {
		for {
			fmt.Println(<-chanel)
		}
	}()
	time.Sleep(time.Millisecond)
}

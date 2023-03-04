package main

import "fmt"

// Fan Out in
func main() {
	g := generateNumbers(8, 10)
	d1 := divideNumbers(g)
	d2 := divideNumbers(g)

	fmt.Println(<-d1)
	fmt.Println(<-d2)
}

func generateNumbers(numbers ...int) chan int { // ele vai retornar um channel que só receb informação
	channel := make(chan int)
	go func() {
		for _, n := range numbers {
			channel <- n
		}
		close(channel)
	}()
	return channel
}

func divideNumbers(input chan int) chan int {
	channel := make(chan int)
	go func() {
		for number := range input {
			channel <- number / 2
		}
		close(channel)
	}()
	return channel
}

package main

import "fmt"

func main() {

	// For
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// While
	x := 1
	for x != 10 {
		fmt.Println(x)
		x++
	}

	for true {
		fmt.Println(x)
		if x == 100 {
			println("CHEGOU")
			break
		}
		x++
	}

	for {
		fmt.Println(x)
		if x == 100000 {
			break
		}
		x++
	}
}

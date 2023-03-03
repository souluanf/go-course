package main

import "fmt"

func main() {
	a := 10
	if a >= 10 {
		fmt.Println("A > 10")
	} else if a < 10 {
		fmt.Println("A > 5")
	}

	b := true

	if x := "Cool"; b {
		fmt.Println(x)
	}
}

package main

import "fmt"

func main() {
	a := "Luan"

	switch a {
	case "Luan":
		fmt.Println("Hey Luan")
	case "Fernandes":
		fmt.Println("Hey Fernandes")
	default:
		fmt.Println("I did not find your name!")
	}
}

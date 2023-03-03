package main

import "fmt"

func main() {
	var x [10]int
	x[0] = 1
	x[1] = 3
	x[2] = 9
	x[3] = 2
	x[4] = 7
	x[5] = 5
	x[6] = 12
	x[7] = 45
	x[8] = 8
	x[9] = 32
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[9])

	var y [10]string
	fmt.Println(y)

	z := [5]int{5, 10, 4, 5, 6}
	fmt.Println(z)
}

package main

import (
	"fmt"
)

func functionName(a int) int {
	return a * a
}

func namedReturn(a string) (x string) {
	x = a
	return
}

func moreReturns(a string, b string) (string, string) {
	return a, b
}

func variadicFunc(x ...int) int { // Diversos valores
	var res int
	for _, v := range x {
		res += v
	}
	return res
}

func functionInsideFunction() func() int {
	x := 10
	return func() int {
		return x * x
	}
}

func main() {
	fmt.Println(functionName(10))

	fmt.Println(namedReturn("Luan"))

	x, y := moreReturns("Luan", "Fernandes")
	fmt.Printf("%s %s\n", x, y)

	result := variadicFunc(10, 20, 30)
	fmt.Println(result)

	// Arrow
	z := 0
	add := func() int {
		z += 2
		return z
	}
	fmt.Println(add())
	fmt.Println(add())

}

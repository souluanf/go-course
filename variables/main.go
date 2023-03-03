package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
)

const (
	Aa string = "Olá"
	Bb        = 66
	Cc        = 77
)

const AA int = 1333

func main() {
	a := 10
	b := "Hello"
	c := 10.45
	d := false
	e := 'W'
	f := `Can use to write long strings`

	const example2 = 10

	fmt.Printf("variable: %v, type: %T \n", a, a)
	fmt.Printf("variable: %v, type: %T \n", b, b)
	fmt.Printf("variable: %v, type: %T \n", c, c)
	fmt.Printf("variable: %v, type: %T \n", d, d)
	fmt.Printf("variable: %v, type: %T \n", e, e)
	fmt.Printf("variable: %v, type: %T \n", f, f)

	u, _ := uuid.NewV4()
	fmt.Printf("hello, %s\n", u)
}

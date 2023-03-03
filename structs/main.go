package main

import "fmt"

type Car struct {
	Name  string
	Year  int
	Color string
}

func (c Car) toString() string {
	return fmt.Sprintf("{\"name\": \"%s\", \"year\": %d, \"color\": \"%s\"}", c.Name, c.Year, c.Color)
}

func main() {
	car1 := Car{"X1", 2016, "Black"}
	car2 := Car{"Corolla", 2015, "Black"}

	fmt.Println(car1.Name, car1.Year, car1.Color)
	fmt.Println(car2.Name, car2.Year, car2.Color)

	fmt.Println(car1.toString())
	fmt.Println(car2.toString())
}

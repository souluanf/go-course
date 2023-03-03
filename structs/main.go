package main

import "fmt"

type Car struct {
	Name  string
	Year  int
	Color string
}

type SuperCar struct {
	Car
	CanFly bool
}

func (c Car) toString() string {
	return fmt.Sprintf("{\"name\": \"%s\", \"year\": %d, \"color\": \"%s\"}", c.Name, c.Year, c.Color)
}

func main() {
	car1 := Car{"X1", 2016, "Black"}
	car2 := Car{"Corolla", 2015, "Black"}
	scar := SuperCar{Car: Car{"Ferrari", 2030, "Red"}, CanFly: false}

	fmt.Println(car1.toString())
	fmt.Println(car2.toString())
	fmt.Println(scar.toString())
}

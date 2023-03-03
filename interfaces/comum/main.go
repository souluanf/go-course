package main

import "fmt"

type Vehicle interface {
	start() string
}

type Car struct {
	name string
}

func (c Car) start() string {
	return "The car " + c.name + " has been started."
}

type Motorcycle struct {
	name string
}

func (m Motorcycle) start() string {
	return "The motorcycle " + m.name + " has been started."
}

func startVehicle(v Vehicle) string {
	return v.start()
}

func main() {

	c := Car{"Gol"}
	m := Motorcycle{"Ninja"}

	fmt.Println(startVehicle(c))
	fmt.Println(startVehicle(m))

	fmt.Println(c.start())
	fmt.Println(m.start())
}

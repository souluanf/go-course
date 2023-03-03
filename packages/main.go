package main

import (
	"fmt"
	"github.com/souluanf/go-course/packages/car"
)

func main() {
	car1 := car.Car{Name: "Gol", Color: "Black"}
	fmt.Println(car1.Start())
}

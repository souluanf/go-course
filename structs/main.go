package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Car struct {
	Name  string `json:"nome"` //Como estará no arquivo json
	Year  int    `json:"-"`    // Usado para nao trazer o valor no Json
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

	car3 := Car{"Gol", 2017, "Yellow"}
	result, _ := json.Marshal(car3)
	fmt.Println(result)
	fmt.Println(bytes.NewBuffer(result))
	fmt.Println(string(result))
}

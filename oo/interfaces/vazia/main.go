package main

import "fmt"

type Names []interface{}

func (n *Names) load() {
	*n = Names{"Luan", "Heimy", "Principe", "Safira"}
}

func (n Names) printNames() {
	fmt.Println(n)
}

func main() {
	var names Names

	names.load()
	names.printNames()
}

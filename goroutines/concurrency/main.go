package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	go runProcess("P1", 20)
	go runProcess("P2", 20)
	go runProcess("P3", 20)
	go runProcess("P4", 20)

	var s string
	_, err := fmt.Scanln(&s)
	if err != nil {
		return
	}
}

func runProcess(name string, total int) {
	for i := 0; i < total; i++ {
		fmt.Println(name, "->", i)
		t := time.Duration(rand.Intn(255))
		time.Sleep(time.Millisecond * t)
	}
	fmt.Println(name, " ganhou!")
	os.Exit(0)
}

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var result int
var m sync.Mutex // Sempre que precisar proteger dados de mudança de valores

func main() {
	go runProcess("P1", 20)
	go runProcess("P2", 20)
	var s string
	_, _ = fmt.Scanln(&s)
	fmt.Println("Final result: ", result)
}

func runProcess(name string, total int) {
	for i := 0; i < total; i++ {
		t := time.Duration(rand.Intn(255))
		time.Sleep(time.Millisecond * t)
		m.Lock()
		result++
		fmt.Println(name, "->", i, " Partial result: ", result)
		m.Unlock()
	}

}

// go run -race main.go

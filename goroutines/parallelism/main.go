package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // Processar com o número máximo de CPus
}

var waitinfGroup sync.WaitGroup

func main() {
	fmt.Println(runtime.NumCPU())
	waitinfGroup.Add(2) // dois processos rodando
	go runProcess("P1", 20)
	go runProcess("P2", 20)
	go runProcess("P3", 20)
	go runProcess("P4", 20)
	go runProcess("P5", 20)
	go runProcess("P6", 20)
	waitinfGroup.Wait()
}

func runProcess(name string, total int) {
	for i := 0; i < total; i++ {
		fmt.Println(name, "->", i)
		t := time.Duration(rand.Intn(255))
		time.Sleep(time.Millisecond * t)
	}
	waitinfGroup.Done()

}

// # Paralelismo
// P1 __________________________________(1 core)
// P2 __________________________________(1 core)

// # Cocurrency (1 core processa tudo)
// P1 ____    _____   ___    ____    __
// P2     ___     ___    ____     ___

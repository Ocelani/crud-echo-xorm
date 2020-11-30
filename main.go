package main

import (
	"fmt"
	"time"

	"github.com/Ocelani/mercafacil/pkg"
	"github.com/Ocelani/mercafacil/pkg/simulation"
)

func main() {
	go pkg.Server()

	time.Sleep(time.Second)
	fmt.Println()

	for i := 5; i >= 0; i-- {
		time.Sleep(time.Second)
		fmt.Printf("\r ••• Simulation starting in: %v", i)
	}

	simulation.Run()
}

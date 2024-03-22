// Package: main Gateway
package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("Orchestrator service is running...")
		time.Sleep(time.Second)
	}
}

// Package: main Gateway
package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("Storage service is running...")
		time.Sleep(time.Second)
	}
}

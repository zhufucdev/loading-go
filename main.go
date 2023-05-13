package main

import (
	"time"
	"fmt"
)

func main() {
	fmt.Print("Loading...")
	for (true) {
		time.Sleep(1 * time.Second)
		fmt.Print(".")
	}
}

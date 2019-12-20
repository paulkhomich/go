package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		for _, r := range "/\\" {
			fmt.Printf("\r%c", r)
			time.Sleep(100 * time.Millisecond)
		}
	}
}

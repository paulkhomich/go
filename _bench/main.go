package main

import (
	"sync"
)

func main() {
	integers := make(chan int)

	go printer(integers)

	integers <- 3
	integers <- 2
	integers <- 6
	close(integers)
}

func printer(s <-chan int) {
	ch := make(chan int)

	go func() {
		var wg sync.WaitGroup

		for i := range s {
			wg.Add(1)
			go func(number int) {
				defer wg.Add(-1)
				ch <- number * number
			}(i)
		}

		go func() {
			wg.Wait()
			close(ch)
		}()
	}()

	for v := range ch {
		println(v)
	}
}

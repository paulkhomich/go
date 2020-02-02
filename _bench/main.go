package main

import (
)

func main() {
	buf := make([]byte, 1024)
	for range []int{1, 2, 3} {
		n, err := os.Stdin.Read(buf)
		if err == io.EOF {
			break
		}
		println(n)
	}
}

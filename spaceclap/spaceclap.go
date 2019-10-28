package main

import "fmt"

func main() {
	bytes := []byte("king  of     kings")

	fmt.Println(bytes)
	bytes = spaceClap(bytes)
	fmt.Println(bytes)
}

func spaceClap(bytes []byte) []byte {
	for i := len(bytes) - 1; i > 0; i-- {
		if bytes[i] == ' ' && bytes[i+1] == ' ' {
			copy(bytes[i:], bytes[i+1:])
			bytes = bytes[:len(bytes)-1]
		}
	}

	return bytes
}

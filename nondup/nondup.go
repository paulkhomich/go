package main

import "fmt"

func main() {
	s := []string{"hello", "world", "freddy", "freddy", "now"}

	fmt.Println(s)
	s = nondup(s)
	fmt.Println(s)
}

func nondup(s []string) []string {
	for i := len(s) - 1; i > 0; i-- {
		if s[i] == s[i - 1] {
			copy(s[i-1:], s[i:])
			s = s[:len(s) - 1]
		}
	}

	return s
}

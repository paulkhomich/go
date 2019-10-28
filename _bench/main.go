package main

import "fmt"

func main() {
	a1 := [...]int{0, 1, 2}
	a2 := a1[:]

	fmt.Println(a2)
}

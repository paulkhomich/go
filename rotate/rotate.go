package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 3, 4, 5}

	rotate(arr[:], 2)
	fmt.Println(arr)
}

func rotate(s []int, n int) {
	temp := append(s[n:], s[:n]...)
	copy(s, temp)
}

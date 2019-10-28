package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 5}

	reverse(&a)
	fmt.Println(a)
}

func reverse(arr *[5]int) {
	for i, j := 0, len(arr)-1; i <= j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

package main

import "sort"

func main() {
	s := []string{"h", "e", "l", "l", "o", "l", "l", "e", "h"}
	sw := []string{"h", "e", "l", "l", "o", "w", "o", "r", "l", "d"}

	println(isPalindrome(sort.StringSlice((s))))
	println(isPalindrome(sort.StringSlice((sw))))
}

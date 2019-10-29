package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	words := map[string]int{}

	file, err := os.Open("text.txt")
	if err != nil {
		fmt.Printf("Err: %s\n", err)
	}

	input := bufio.NewScanner(file)
	input.Split(bufio.ScanWords)

	for input.Scan() {
		word := input.Text()
		words[word]++
	}
	if err := input.Err(); err != nil {
		fmt.Printf("Err: %s\n", err)
	}
	file.Close()

	fmt.Printf("word\t\tcount\n")
	for k, v := range words {
		fmt.Printf("%s\t\t%d\n", k, v)
	}
}

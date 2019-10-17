package main

import (
	"os"
	"fmt"
	"bufio"
)

func main() {
	args := os.Args[1:];
	if len(args) == 0 {
		return
	}

	counts := make(map[string]int)
	files := make(map[string]map[string]bool)

	for _, filename := range args {
		file, err := os.Open(filename);
		if err != nil {
			fmt.Fprintf(os.Stderr, "[error] %v", err)
			continue
		}

		countLines(file, counts, files)
		file.Close()
	}

	for line, n := range counts {
		if n > 1 {
			for filename := range files[line] {
				fmt.Printf("%s ", filename);
			}

			fmt.Println(line);
		}
	}
}

func countLines(file *os.File, counts map[string]int, files map[string]map[string]bool) {
	input := bufio.NewScanner(file)

	for input.Scan() {
		text := input.Text()

		if files[text] == nil {
			files[text] = make(map[string]bool)
		}

		counts[text]++
		files[text][file.Name()] = true
	}
}
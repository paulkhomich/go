package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	unitypes := map[string]int{}

	reader := bufio.NewReader(os.Stdin)
	for {
		r, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Err: %s", err)
		}

		if unicode.IsLetter(r) {
			unitypes["letter"]++
		} else if unicode.IsNumber(r) {
			unitypes["number"]++
		} else if unicode.IsMark(r) {
			unitypes["mark"]++
		} else if unicode.IsPunct(r) {
			unitypes["punct"]++
		} else if unicode.IsSpace(r) {
			unitypes["space"]++
		} else {
			unitypes["other"]++
		}
	}

	fmt.Printf("type\tcount\n")
	for i, v := range unitypes {
		fmt.Printf("%s\t%d\n", i, v)
	}
}

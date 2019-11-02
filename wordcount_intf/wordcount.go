package main

import (
	"fmt"
	"unicode"
)

// WordCount is int-counter for words in input stream
type WordCount int

func (counter *WordCount) Write(bytes []byte) (int, error) {
	trig := true
	for _, v := range bytes {
		if unicode.IsSpace(rune(v)) {
			trig = true
		} else if trig {
			trig = false
			(*counter)++
		}
	}

	return len(bytes), nil
}

func main() {
	var wc WordCount
	wc.Write([]byte("Hello man"))
	println(wc)

	fmt.Fprintf(&wc, "Either or neither")
	println(wc)
}

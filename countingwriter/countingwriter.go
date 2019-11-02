package main

import (
	"fmt"
	"os"
)

func main() {
	w := os.Stdout
	wc, counter := countingWirter(w)

	fmt.Fprintf(wc, "hello!\n")
	println(*counter)
}

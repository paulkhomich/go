package main

import (
	"fmt"
	"os"
)

func main() {
	r := os.Stdin
	lr := limitReader(r, 20)

	buf := make([]byte, 12)

	n, err := lr.Read(buf)
	fmt.Println(buf, n, err)
	n, err = lr.Read(buf)
	fmt.Println(buf, n, err)
}

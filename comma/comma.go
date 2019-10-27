package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("12345678"))
}

// Comma format raw numbers input string to [1,234,567] format
func comma(s string) string {
	var buf bytes.Buffer

	lenMod := len(s) % 3
	for i, v := range s {
		if i > 0 && (i-lenMod)%3 == 0 {
			buf.WriteByte(',')
		}

		buf.WriteByte(byte(v))
	}

	return buf.String()
}

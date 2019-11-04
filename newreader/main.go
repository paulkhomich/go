package main

import "fmt"

func main() {
	reader := newReader("Hello")

	buf := make([]byte, 3, 3)
	n, err := reader.Read(buf)
	fmt.Println(buf)
	fmt.Println(n)
	fmt.Println(err)
	n, err = reader.Read(buf)
	fmt.Println(buf)
	fmt.Println(n)
	fmt.Println(err)
}

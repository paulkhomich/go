package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	num := make(chan string)

	go hi(num)

	fmt.Println(<-num)
}

func hi(ch chan string) {
	r := bufio.NewScanner(os.Stdin)
	r.Scan()
	ch <-r.Text()
}

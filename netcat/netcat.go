package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func handle(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%w", err)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("usage: netcat \"addr:port\"\n")
	}
	conn, err := net.Dial("tcp", os.Args[1])
	handle(err)

	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		done <- struct{}{}
	}()

	io.Copy(conn, os.Stdin)
	conn.Close()
	<-done
}

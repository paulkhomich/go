package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	listener, _ := net.Listen("tcp", "localhost:8080")

	for {
		conn, _ := listener.Accept()
		handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	println(conn.RemoteAddr().String())
	ticker := time.NewTicker(2 * time.Second)

	go readConn(conn)

	for {
		<-ticker.C
		io.WriteString(conn, "echo!")
		println("send echo...")
	}
}

func readConn(conn net.Conn) {
	remote := bufio.NewScanner(conn)
	remote.Split(bufio.ScanRunes)
	for remote.Scan() {
		fmt.Printf(remote.Text())
	}
}
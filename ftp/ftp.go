package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			continue
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)

	for {
		io.WriteString(conn, fmt.Sprintf("%s:~ %s$ ", conn.LocalAddr(), conn.RemoteAddr()))

		str, _ := r.ReadString('\n')
		cmd := strings.Split(strings.Trim(str, " \n"), " ")

		switch cmd[0] {
		case "":
			io.WriteString(conn, "")
		case "ls":
			{
				out, err := exec.Command("ls").Output()
				if err != nil {
					fmt.Fprintf(os.Stderr, "%s\n", err)
					continue
				}

				conn.Write(out)
			}
		case "get":
			{
				out, err := exec.Command("cat", cmd[1]).Output()
				if err != nil {
					io.WriteString(conn, err.Error() + "\n")
					continue
				}

				conn.Write(out)
			}
		case "close":
			{
				io.WriteString(conn, "bye!\n")
				return
			}
		default:
			io.WriteString(conn, "unknown command\n")
		}

	}
}

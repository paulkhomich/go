package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

type msg struct {
	user *user
	text string
}

type user struct {
	ip       string
	username string
	ch       chan msg
}

var registration = make(chan user)
var messages = make(chan msg)
var quit = make(chan user)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)

		os.Exit(1)
	}

	go dispenser(listener.Addr().String())

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Обработка подключения
		go handle(conn)
	}
}

func dispenser(adr string) {
	online := make(map[user]bool)
	server := user{
		adr,
		"server",
		nil,
	}

	sender := func(message msg) {
		for user := range online {
			user.ch <- message
		}
	}

	for {
		select {
		case user := <-registration:
			online[user] = true
			sender(msg{&server, user.username + " подключился"})
		case message := <-messages:
			sender(message)
		case user := <-quit:
			delete(online, user)
			sender(msg{&server, user.username + " ушел"})
			close(user.ch)
		}
	}
}

func handle(conn net.Conn) {
	username := greeting(conn)
	ch := make(chan msg)
	who := user{
		conn.RemoteAddr().String(),
		username,
		ch,
	}

	registration <- who

	go printer(conn, ch)

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		messages <- msg{
			&who,
			scanner.Text(),
		}
	}

	quit <- who
	conn.Close()
}

func greeting(conn net.Conn) string {
	io.WriteString(conn, "Welcome!\n")

	username := ""
	scanner := bufio.NewScanner(conn)
	for len(username) == 0 {
		io.WriteString(conn, "Enter your username: ")
		scanner.Scan()
		println("got")
		username = scanner.Text()
	}

	return username
}

func printer(conn net.Conn, ch <-chan msg) {
	for m := range ch {
		if m.user.ch == nil {
			fmt.Fprintln(conn, "\u001b[93m"+m.text+"\u001b[39m")
		} else if m.user.ch == ch {
			fmt.Fprintln(conn, "[\u001b[36m"+m.user.username+"\u001b[39m]: "+m.text)
		} else {
			fmt.Fprintln(conn, "[\u001b[31m"+m.user.username+"\u001b[39m]: "+m.text)
		}
	}
}

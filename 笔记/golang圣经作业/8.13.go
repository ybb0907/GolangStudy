package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

type client chan<- string // an outgoing message channel

type user struct {
	ch   client
	name string
}

var (
	entering = make(chan user)
	leaving  = make(chan user)
	messages = make(chan string, 1) // all incoming client messages
)

func broadcaster() {
	clients := make(map[user]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli.ch <- msg
			}
		case cli := <-entering:
			clients[cli] = true
			msg := "now , this is : "
			for client := range clients {
				msg += client.name + ", "
			}
			messages <- msg

		case cli := <-leaving:
			msg := "this conn is leaving : " + cli.name
			messages <- msg
			delete(clients, cli)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	newClinet := user{ch: ch, name: who}
	entering <- newClinet

	t := time.NewTicker(2 * time.Second)

	isClose := make(chan string, 1) // outgoing client messages

	go func() {
		for {
			select {
			case _ = <-t.C:
				isClose <- "1"
			}
		}
	}()

	go func() {
		for {
			select {
			case _ = <-isClose:
				messages <- who + " has left"

				// time.Sleep(time.Second)

				leaving <- user{
					ch:   ch,
					name: who,
				}
				conn.Close()
			}
		}
	}()

	input := bufio.NewScanner(conn)
	for input.Scan() {
		// t.Reset(2 * time.Second)
		messages <- who + ": " + input.Text()
	}
	// NOTE: ignoring potential errors from input.Err()

	isClose <- "1"
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

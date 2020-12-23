package main

import (
	"fmt"
	"net"
	"bufio"
)

var clients = []client{}

func main() {
	dstream, err := net.Listen("tcp", ":8000")
	defer dstream.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	idCounter := 1

	for {
		conn, err := dstream.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("New connection detected!")
		c := client{conn, idCounter}
		clients = append(clients, c)

		idCounter++
		go c.handle()
	}
}

func (currClient client) handle() {
	for {
		data, err := bufio.NewReader(currClient.conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, client := range clients {
			if (client.id != currClient.id) {
				s := fmt.Sprintf("User %d: %s\n", currClient.id, data)
				client.conn.Write([]byte(s))
				fmt.Println(s)
			}
		}
		// fmt.Println(data)
	}

	currClient.conn.Close()
}

type client struct {
	conn net.Conn
	id int
}
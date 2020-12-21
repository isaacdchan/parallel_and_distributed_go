package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
)

var conns = map[net.Addr]net.Conn{}

func main() {
	port := flag.String("port number", "8000", "Please specify the port number")
	numClients := flag.Int("number of clients", 2, "Please specify the number of clients who will participate")
	dstream, err := net.Listen("tcp", ":"+*port)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer dstream.Close()
	conns = make(map[net.Addr]net.Conn, *numClients)

	for {
		conn, err := dstream.Accept()
		
		if err != nil {
			fmt.Println(err)
			return
		}

		defer conn.Close()
		addr := conn.RemoteAddr()
		conns[addr] = conn

		go handle(conn)
	}
}

func handle(client net.Conn) {
	for {
		fmt.Println("Reading")
		data, _, err := bufio.NewReader(client).ReadLine()
		fmt.Println("Done Reading")
	
		if err != nil {
			fmt.Println(err)
			return
		}
	
		s := string(data)
		s += "\n"
		fmt.Println("String: " + s)
	
		for _, conn := range conns {
			fmt.Println("Writing!")
			conn.Write([]byte(s))
			fmt.Println("Done Writing!")
		}
		fmt.Println(4)
	}
}
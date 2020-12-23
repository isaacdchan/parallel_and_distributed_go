package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
)

var conns = map[string]net.Conn{}

func main() {
	port := flag.String("port number", "8000", "Please specify the port number")
	numClients := flag.Int("number of clients", 2, "Please specify the number of clients who will participate")
	dstream, err := net.Listen("tcp", ":"+*port)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer dstream.Close()
	conns = make(map[string]net.Conn, *numClients)

	for {
		conn, err := dstream.Accept()
		
		if err != nil {
			fmt.Println(err)
			return
		}

		defer conn.Close()
		addr := conn.RemoteAddr().String()
		conns[addr] = conn

		go handle(conn)
	}
}

func handle(client net.Conn) {
	for {
		data, _, err := bufio.NewReader(client).ReadLine()
	
		if err != nil {
			fmt.Println(err)
			return
		}
	
		msg := string(data)
		addr := client.RemoteAddr().String()
		outgoingMsg := []byte(addr + ": " + msg + "\n")
	
		for connAddr, conn := range conns {
			if connAddr != addr {
				conn.Write(outgoingMsg)
			}
		}
	}
}
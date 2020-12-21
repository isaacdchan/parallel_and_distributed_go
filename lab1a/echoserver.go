package main

import (
	"fmt"
	"net"
	"bufio"
)

func main() {
	dstream, err := net.Listen("tcp", ":8000")
	defer dstream.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		conn, err := dstream.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {

	data, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}

	conn.Write([]byte(data))
	conn.Close()
}

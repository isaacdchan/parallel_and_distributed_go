package main

import (
	"CSS434/lab1b/utils"
	"bufio"
	"flag"
	"fmt"
	"net"
)

var multiplier = 2

func main() {
	port := flag.String("port number", "8000", "Please specify the port number")
	dstream, err := net.Listen("tcp", ":" + *port)
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

		fmt.Println("Connection Accepted!")
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	fmt.Println("Reading!")
	// data, err := ioutil.ReadAll(conn)
	data, err := bufio.NewReader(conn).ReadBytes('\n')

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Finished Reading!")

	s := string(data)
	s = s[:len(s) - 1]
	ia := utils.StringToIntArray(s)

	for i := range ia {
		ia[i]*=multiplier
	}
	multiplier *= 2

	ba := utils.IntArrayToByteArray(ia)

	fmt.Println("Writing")
	conn.Write(ba)
	fmt.Println("Finished writing")
	conn.Close()
}

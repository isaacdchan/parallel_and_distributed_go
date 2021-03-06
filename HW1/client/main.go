package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"sync"
)

func main() {
	port := flag.String("port number", "8000", "Please specify the port number")

	serv, err := net.Dial("tcp", ":" + *port)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer serv.Close()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		handleOutgoing(serv)
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		handleIncoming(serv)
		wg.Done()
	}()

	wg.Wait()
}

func handleOutgoing(serv net.Conn) {
	for {
		reader := bufio.NewReader(os.Stdin)
		// fmt.Print("Enter text: ")
		s, _ := reader.ReadString('\n')
		ba := []byte(s)

		serv.Write(ba)
	}
}

func handleIncoming(serv net.Conn) {
	for {
		data, _, err := bufio.NewReader(serv).ReadLine()

		if err != nil {
			fmt.Println(err)
			return
		}

		msg := string(data)
		fmt.Println("\n" + msg)
	}
}
package main

import (
	"CSS434/lab1b/utils"
	"bufio"
	"flag"
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {
	port := flag.String("port number", "8000", "Please specify the port number")
	arrayLength := flag.Int("array length", 10, "Please specify the array length")
	fmt.Println("Connecting!")
	// no input validation yet

	conn, err := net.Dial("tcp", ":" + *port)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Connected!")

	// send the array to server to be processed
	sendInitialArray(conn, *arrayLength)

	// wait for the server to process
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		handle(conn)
		wg.Done()
	}()

	wg.Wait()
}

func sendInitialArray(conn net.Conn, arrayLength int) {
	ia := make([]int, arrayLength)
	for i := range ia {
		ia[i] = i + 1
	}
	ba := utils.IntArrayToByteArray(ia)

	fmt.Println("Writing")
	conn.Write(ba)
	fmt.Println("Finished writing")
}

func handle(conn net.Conn) {
	time.Sleep(time.Second * 3)
	fmt.Println("Reading!")
	data, err := bufio.NewReader(conn).ReadBytes('\n')

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Finished Reading!")

	s := string(data)
	// remove trailing newline rune
	s = s[:len(s) - 1]
	ia := utils.StringToIntArray(s)
	fmt.Println(ia)
}
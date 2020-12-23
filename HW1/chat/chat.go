package main

import (
	"CSS434/HW1/utils"
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
)

type chat struct {
	port       int
	hosts      []string
	numHosts   int
	rank       int
	stamps     []int
	conns      []net.Conn
	msgQueue   []string
	stampQueue [][]int
	srcQueue   []int
}

func (c chat) establishConnections() {
	dstream, err := net.Listen("tcp", ":" + strconv.Itoa(c.port))
	if err != nil {
		fmt.Println(err)
		return
	}

	defer dstream.Close()
	c.acceptConnections(dstream)
	c.sendConnections(dstream)

}

func (c chat) acceptConnections(dstream net.Listener) {
	for i := c.numHosts; i > c.rank; i++ {
		conn, err := dstream.Accept()

		if err != nil {
			fmt.Println(err)
			return
		}

		defer conn.Close()
		srcAddr := conn.RemoteAddr().String()
		srcIndex := -1
		
		fmt.Println("Accepted connection from " + srcAddr)
		for i, hostName := range c.hosts {
			if (hostName == srcAddr) {
				srcIndex = i
			}
		}

		c.conns[srcIndex] = conn
	}
}

func (c chat) sendConnections(dstream net.Listener) {
	for i:= 0; i < c.rank; i++ {
		destAddr := c.hosts[i] + ":" + strconv.Itoa(c.port)
		dest, err := net.Dial("tcp", destAddr)

		if err != nil {
			fmt.Println(err)
			return
		}

		c.conns[i] = dest
	}
}

func (c chat) determineRank(hosts []string) {
	c.rank = -1

	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr).String()

	for i, host := range hosts {
		if (host == localAddr) {
			c.rank = i
		}
	}
}

func (c chat) handleOutgoing() {
	for {
		for _, conn := range c.conns {
			reader := bufio.NewReader(os.Stdin)
			// fmt.Print("Enter text: ")
			s, _ := reader.ReadString('\n')
			ba := []byte(s)

			c.handleOutgoingStamps(conn)
			conn.Write(ba)
		}
	}
}
func (c chat) handleOutgoingStamps(conn net.Conn) {
	c.stamps[c.rank]++
	ba := utils.IntArrayToByteArray(c.stamps)
	conn.Write(ba)
}

func (c chat) handleIncoming() {
	for i, conn := range c.conns {
		for {
			stampData, _, err := bufio.NewReader(conn).ReadLine()
			msgData, _, err := bufio.NewReader(conn).ReadLine()

			if err != nil {
				fmt.Println(err)
				return
			}

			msg := string(msgData)
			srcStamps := utils.ByteArrayToIntArray(stampData)

			if (!c.canAcceptMsg(i, srcStamps)) {
				c.msgQueue = append(c.msgQueue, msg)
				c.stampQueue = append(c.stampQueue, srcStamps)
				c.srcQueue = append(c.srcQueue, i)
			} else {
				fmt.Println(c.hosts[i] + ": " + msg)
				c.stamps[i]++
			}
		}
	}
}

func (c chat) canAcceptMsg(src int, srcStamps []int) (bool) {
	for i, stamp := range c.stamps {
		srcStamp := srcStamps[i]

		// if src has an additional msg to curr pending
		if (i == src && stamp + 1 != srcStamp) {
			return false
		} else {
			// src is too far ahead regarding communications between src and other chats
			if (srcStamp > stamp) {
				return false
			}
		}
	}

	return true
}

func (c chat) handleStampQueue() {
	for {
		for i, pendingStamp := range c.stampQueue {
			pendingSrc := c.srcQueue[i]
			pendingMsg := c.msgQueue[i]

			if (c.canAcceptMsg(pendingSrc, pendingStamp)) {
				c.stampQueue = append(c.stampQueue[:i], c.stampQueue[i+1:]...)
				c.srcQueue = append(c.srcQueue[:i], c.srcQueue[i+1:]...)
				c.msgQueue = append(c.msgQueue[:i], c.msgQueue[i+1:]...)

				// record in stamps that pendingMsg was accepted
				c.stamps[pendingSrc]++
				fmt.Println(c.hosts[i] + ": " + pendingMsg)
			}
		}
	}

}
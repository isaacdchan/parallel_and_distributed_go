package main

import (
	"flag"
	"net"
	"sync"
)


var hosts []string

func main() {
	port := flag.Int("port number", 0, "please enter the port number")
	hosts := flag.Args()
	hosts = []string{"172.24.186.253.34285", }

	c := newChat(*port, hosts)
	c.establishConnections()

	var wg sync.WaitGroup

	wg.Add(3)
	go func() {
		c.handleOutgoing()
		wg.Done()
	}()
	go func() {
		c.handleIncoming()
		wg.Done()
	}()
	go func() {
		c.handleStampQueue()
		wg.Done()
	}()

	wg.Wait()
}

func newChat(port int, hosts []string) *chat {
	c := new(chat)

	c.port = port
	c.hosts = hosts
	numHosts := len(hosts)

	c.numHosts = numHosts
	c.stamps = make([]int, numHosts)
	c.conns = make([]net.Conn, numHosts)
	c.msgQueue = []string{}
	c.stampQueue = [][]int{}
	c.srcQueue = []int{}
	c.determineRank(hosts)

	return c
}
package main

import (
	"net"
	"log"
	"flag"
	"strings"
	"encoding/binary"
)

func main() {
	var nodelay bool
	var msgSize int
	var n int
	flag.BoolVar(&nodelay, "nodelay", false, "")
	flag.IntVar(&msgSize, "size", 128, "")
	flag.IntVar(&n, "n", 1000000, "")
	flag.Parse()

	log.Printf("nodelay: %v, size: %d, n: %d\n", nodelay, msgSize, n)

	c, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		log.Fatalln(err)
	}

	c.(*net.TCPConn).SetNoDelay(nodelay)
	msg := []byte(strings.Repeat("A", msgSize))
	end := []byte("end")

	for i := 0; i < n; i++ {
		binary.Write(c, binary.BigEndian, uint32(len(msg)))
		c.Write(msg)
	}

	binary.Write(c, binary.BigEndian, uint32(len(end)))
	c.Write(end)
}

package main

import (
	"encoding/binary"
	"flag"
	"log"
	"net"
	"strings"
)

func main() {
	var nodelay bool
	var msgSize int
	var n int
	flag.BoolVar(&nodelay, "nodelay", false, "")
	flag.IntVar(&msgSize, "size", 128, "")
	flag.IntVar(&n, "n", 1000000, "")
	flag.Parse()

	var nodelayState string
	if nodelay {
		nodelayState = ", nodelay"
	}

	log.Printf("size: %d, n: %d%s\n", msgSize, n, nodelayState)

	c, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		log.Fatalln(err)
	}

	c.(*net.TCPConn).SetNoDelay(nodelay)
	payload := []byte(strings.Repeat("A", msgSize))
	end := []byte("end")

	for i := 0; i < n; i++ {
		msg := make([]byte, 4 + len(payload))
		binary.BigEndian.PutUint32(msg, uint32(len(payload)))
		copy(msg[4:], payload)
		c.Write(msg)
	}

	binary.Write(c, binary.BigEndian, uint32(len(end)))
	c.Write(end)
}

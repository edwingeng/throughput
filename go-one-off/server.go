package main

import (
	"encoding/binary"
	"flag"
	"io"
	"log"
	"net"
	"sync"
	"time"
)

func main() {
	var n int
	flag.IntVar(&n, "n", 1, "")
	flag.Parse()
	if n <= 0 {
		n = 1
	}

	addr := ":8888"
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln(err)
	}

	defer l.Close()
	log.Println("listening on", addr)

	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		c, err := l.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		wg.Add(1)
		go handle(c, &wg)
	}

	wg.Wait()
}

func handle(c net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	defer c.Close()

	header := make([]byte, 4)
	var numMsg int
	var start time.Time
	var err error

	for {
		_, err = io.ReadFull(c, header)
		if err != nil {
			log.Fatalln(err)
		}

		size := binary.BigEndian.Uint32(header)
		if size > 1024*64 {
			log.Fatalln("payload is too big")
		}

		if size > 0 {
			payload := make([]byte, size)
			_, err = io.ReadFull(c, payload)
			if err != nil {
				log.Fatalln(err)
			}
			if size == 3 && string(payload) == "end" {
				break
			}
		}

		numMsg++
		if numMsg == 1 {
			start = time.Now()
		}
	}

	dt := time.Now().Sub(start).Seconds()
	throughput := int(float64(numMsg) / dt)
	log.Printf("n: %d, time: %.3f, throughput: %d\n", numMsg, dt, throughput)
}

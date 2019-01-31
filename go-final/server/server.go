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

	var start time.Time
	buf := make([]byte, 4096)
	var bufLen int
	var numMsg int

	for {
		bytesRead, err := io.ReadAtLeast(c, buf[bufLen:], 1)
		bufLen += bytesRead
		var off int

		for bufLen-off >= 4 {
			size := binary.BigEndian.Uint32(buf[off:])
			if size > 1024*64 {
				log.Fatalln("payload is too big")
			}
			fullMsgSize := 4 + int(size)
			if off+fullMsgSize > bufLen {
				break
			}

			if size > 0 {
				payload := make([]byte, size)
				copy(payload, buf[off+4:off+fullMsgSize])
				if size == 3 && string(payload) == "end" {
					break
				}
			}
			off += fullMsgSize

			numMsg++
			if numMsg == 1 {
				start = time.Now()
			}
		}
		if off < bufLen && off > 0 {
			copy(buf[:bufLen-off], buf[off:])
		}
		if off > 0 {
			bufLen -= off
		}
		if err != nil {
			if err == io.EOF || err == io.ErrUnexpectedEOF {
				break
			} else {
				log.Fatalln(err)
			}
		}
	}

	dt := time.Now().Sub(start).Seconds()
	throughput := int(float64(numMsg) / dt)
	log.Printf("n: %d, time: %.3f, throughput: %d\n", numMsg, dt, throughput)
}

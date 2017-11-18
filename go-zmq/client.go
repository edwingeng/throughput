package main

import (
	"flag"
	"log"
	"strings"
	zmq "github.com/zeromq/goczmq"
	"time"
)

func main() {
	var msgSize int
	var n int
	flag.IntVar(&msgSize, "size", 128, "")
	flag.IntVar(&n, "n", 1000000, "")
	flag.Parse()

	log.Printf("size: %d, n: %d, nodelay\n", msgSize, n)

	addr := "tcp://127.0.0.1:8888"
	dealer, err := zmq.NewDealer(addr)
	if err != nil {
		log.Fatalln(err)
	}

	defer dealer.Destroy()
	payload := []byte(strings.Repeat("A", msgSize))
	end := []byte("end")

	for i := 0; i < n; i++ {
		err = dealer.SendFrame(payload, zmq.FlagNone)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = dealer.SendFrame(end, zmq.FlagNone)
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Second * 3)
}

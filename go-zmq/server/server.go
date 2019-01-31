package main

import (
	"log"
	"time"

	zmq "github.com/zeromq/goczmq"
)

func main() {
	addr := "tcp://*:8888"
	router, err := zmq.NewRouter(addr)
	if err != nil {
		log.Fatalln(err)
	}

	defer router.Destroy()
	log.Println("listening on", addr)

	var numMsg int
	var start time.Time
	var even = true

loop:
	for {
		frames, err := router.RecvMessage()
		if err != nil {
			log.Fatalln(err)
		}

		for i := 0; i < len(frames); i++ {
			even = !even
			if even {
				if len(frames[i]) > 0 && string(frames[i]) == "end" {
					break loop
				}
				numMsg++
				if numMsg == 1 {
					start = time.Now()
				}
			}
		}
	}

	dt := time.Now().Sub(start).Seconds()
	throughput := int(float64(numMsg) / dt)
	log.Printf("n: %d, time: %.3f, throughput: %d\n", numMsg, dt, throughput)
}

package main

import (
	"bufio"
	"log"
	"net/http"
	"time"
)

func main() {
	log.Println("Tester is live")

	for i := 0; i < 2040; i++ {
		time.Sleep(time.Millisecond * 10)
		log.Println("count:", i)
		go Stress()
	}

	for {
		time.Sleep(time.Second * 5)
		log.Println("still alive!")
	}
}

func Stress() {
	resp, err := http.Get("http://raspberrypi4.local:9999/sse2")
	// resp, err := http.Get("http://localhost:9999/sse2")
	if err != nil {
		log.Println("cloud not connect", err)
	}

	if resp != nil && resp.Body != nil {
		reader := bufio.NewReader(resp.Body)
		for {
			_, err := reader.ReadBytes('\n')
			if err != nil {
				log.Println("00")
			}

			log.Println(".")
		}
	}

	log.Println("BYE!!")
}

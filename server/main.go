package main

import (
	"fmt"
	"net"
	"time"

	"github.com/google/uuid"
)

func _setInterval(c net.Conn, interval time.Duration) chan<- bool {
	ticker := time.NewTicker(interval)
	stopIt := make(chan bool)
	go func() {
		for {
			select {
			case <-stopIt:
				fmt.Println("stop setInterval")
				return
			case <-ticker.C:
				sendRequestDataCommand(c)
			}
		}

	}()

	return stopIt
}

func createNewSession(c net.Conn) {
	id := uuid.New()
	// TODO: создавать в хранилище запись <id> - <status: ACTIVE>

	fmt.Printf("Created new session: %s\t\t%s\n", id, c.RemoteAddr().String())
}

func sendInitialMarkup(c net.Conn) {
	fmt.Println("Initial markup was sent!")
	c.Write([]byte(`{"id":1,"data":["a","b"]}`))
}

func sendRequestDataCommand(c net.Conn) {
	fmt.Println("REQUEST_DATA was sent!")
	c.Write([]byte(`REQUEST_DATA`))
}

func HandleConnection(c net.Conn) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	createNewSession(c)
	sendInitialMarkup(c)

	go _setInterval(c, 5*time.Second)

	fmt.Printf("Waiting for data from %s\n", c.RemoteAddr().String())

	for {
		// Make a buffer to hold incoming data.
		buf := make([]byte, 1024)
		// Read the incoming connection into the buffer.
		_, err := c.Read(buf)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
		}

		fmt.Printf("Received from %s: %s\n", c.RemoteAddr().String(), buf)
	}

	// c.Close()
}

func main() {
	fmt.Println("Launching server...")

	// Устанавливаем прослушивание порта
	ln, _ := net.Listen("tcp", ":8081")

	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go HandleConnection(c)
	}
}

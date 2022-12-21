package main

import (
	"encoding/json"
	"fmt"
	"net"
	"strings"
	"time"
)

var markup Markup

func scheduleDataRequests(c net.Conn, interval time.Duration) chan<- bool {
	ticker := time.NewTicker(interval)
	stopIt := make(chan bool)
	go func() {
		for {
			select {
			case <-stopIt:
				fmt.Println("Stopping data requests scheduling...")
				return
			case <-ticker.C:
				sendCommand(c, REQUEST_DATA)
			}
		}

	}()

	return stopIt
}

func sendInitialMarkup(c net.Conn) {
	fmt.Println("Initial markup was sent!")
	c.Write([]byte(`{"id":1,"data":["a","b"]}`))
}

// func sendCommand(c net.Conn, command Command) {
func sendCommand(c net.Conn, command string) {
	c.Write([]byte(command))
	fmt.Printf("%s was sent!\n", command)
}

func handleData(c net.Conn, data []byte) {
	if isJson(data) {
		json.Unmarshal(data, &markup)
	} else {
		command := strings.TrimSpace(string(data))

		switch command {
		case "OPEN_SESSION":
			{
				OpenSession(c)
				go scheduleDataRequests(c, 5*time.Second)
				sendInitialMarkup(c)
			}
		case "PING":
			{
				sendCommand(c, OK)
			}
		case "CLOSE_SESSION":
			{
				CloseSession(c)
				// go scheduleDataRequests(c, 5*time.Second) // TODO: disable
			}
		default:
			{
				fmt.Printf("[%s] Command '%s' is not supported\n", c.RemoteAddr().String(), command)
			}
		}
	}
}

func HandleConnection(c net.Conn) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	// fmt.Printf("Waiting for data from %s\n", c.RemoteAddr().String()) // debug
	for {
		buf := make([]byte, 1024)
		_, err := c.Read(buf)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			break
		}

		fmt.Println()
		fmt.Printf("[%s] Received: %s\n", c.RemoteAddr().String(), buf)
		handleData(c, buf)
	}

	c.Close()
}

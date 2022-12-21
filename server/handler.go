package main

import (
	"encoding/json"
	"fmt"
	"net"
	"time"
)

var markup Markup

func scheduleDataRequests(c net.Conn, interval time.Duration) {
	stopIt := make(chan bool)
	stoppers[c.LocalAddr().String()] = stopIt

	ticker := time.NewTicker(interval)
	go func() {
		for {
			select {
			case <-stopIt:
				fmt.Printf("[%s] Stopping data requests scheduling...\n", c.RemoteAddr().String())
				return
			case <-ticker.C:
				sendCommand(c, REQUEST_DATA)
			}
		}

	}()
}

func sendInitialMarkup(c net.Conn) {
	c.Write([]byte(`{"id":1,"data":["a","b"]}`))
	fmt.Printf("[%s] Initial markup was sent!\n", c.RemoteAddr().String())
}

// func sendCommand(c net.Conn, command Command) {
func sendCommand(c net.Conn, command string) {
	c.Write([]byte(command))
	fmt.Printf("[%s] Command '%s' was sent!\n", c.RemoteAddr().String(), command)
}

func handleData(c net.Conn, buf []byte, n int) {
	data := buf[0:n]

	if isJson(data) {
		json.Unmarshal(data, &markup)
	} else {
		command := string(data)

		switch command {
		case "OPEN_SESSION":
			OpenSession(c)
			go scheduleDataRequests(c, 5*time.Second)
			sendInitialMarkup(c)
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
	fmt.Printf("[%s] TCP connection established\n", c.RemoteAddr().String())
	for {
		buf := make([]byte, 1024)
		n, err := c.Read(buf)
		if err != nil {
			fmt.Printf("[%s] Error reading: %s\n", c.RemoteAddr().String(), err.Error())
			CloseSession(c)
			fmt.Printf("[%s] TCP connection closed\n", c.RemoteAddr().String())
			c.Close()
			break
		}

		fmt.Println()
		fmt.Printf("[%s] Received: %s\n", c.RemoteAddr().String(), buf)
		handleData(c, buf, n)
	}

	c.Close()
}

package main

import (
	"fmt"
	"net"
)

type SessionStatus string

const (
	ACTIVE SessionStatus = "ACTIVE"
	CLOSED SessionStatus = "CLOSED"
)

type Session struct {
	Addr   string
	Status SessionStatus
}

var sessions map[string]Session = make(map[string]Session)
var stoppers map[string]chan bool = make(map[string]chan bool)

func OpenSession(c net.Conn) {
	addr := c.RemoteAddr().String()
	sessions[addr] = Session{
		Addr:   addr,
		Status: ACTIVE,
	}

	fmt.Printf("[%s] Opened new session\n", addr)
}

func CloseSession(c net.Conn) {
	stoppers[c.LocalAddr().String()] <- true

	addr := c.RemoteAddr().String()
	sessions[addr] = Session{
		Addr:   addr,
		Status: CLOSED,
	}

	fmt.Printf("[%s] Session closed\n", addr)
}
